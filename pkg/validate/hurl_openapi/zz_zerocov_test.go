//ff:func feature=validate type=test topic=hurl-openapi
//ff:what zz_zerocov_test — hurl_openapi 0% 함수 (collectCoveredStatusCodes / xoh12CheckRoute / xoh13CheckFunc / xoh13CheckGuard / xoh13CheckHappyPath) 단위 테스트
package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestCollectCoveredStatusCodes_ZeroCov(t *testing.T) {
	routes := []apiRoute{{
		Path:     "/users",
		Method:   "GET",
		Segments: []string{"users"},
		Op:       &openapi3.Operation{OperationID: "ListUsers"},
	}}
	entries := []hurl.HurlEntry{
		{Method: "GET", Path: "/users", StatusCode: "200"},
		{Method: "GET", Path: "/users", StatusCode: "401"},
		{Method: "GET", Path: "/users", StatusCode: ""},        // skipped: no code
		{Method: "GET", Path: "/unknown", StatusCode: "200"},   // skipped: no route
	}
	covered := collectCoveredStatusCodes(entries, routes)
	set := covered["ListUsers"]
	if !set["200"] || !set["401"] {
		t.Fatalf("expected 200 & 401 covered, got %v", set)
	}
	if len(set) != 2 {
		t.Errorf("expected exactly 2 codes, got %v", set)
	}
}

func TestXoh12CheckRoute_ZeroCov(t *testing.T) {
	route := apiRoute{
		Op:        &openapi3.Operation{OperationID: "ListUsers"},
		Responses: map[string]bool{"200": true, "401": true, "500": true},
	}
	covered := map[string]map[string]bool{"ListUsers": {"200": true}}
	// 401 declared but uncovered (500 excluded) → 1 warning.
	diags := xoh12CheckRoute(route, covered)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d: %+v", len(diags), diags)
	}

	// nil op → nil.
	if got := xoh12CheckRoute(apiRoute{}, covered); got != nil {
		t.Error("nil op should return nil")
	}
	// all declared covered → nil.
	full := map[string]map[string]bool{"ListUsers": {"200": true, "401": true}}
	if got := xoh12CheckRoute(route, full); got != nil {
		t.Errorf("fully covered should return nil, got %v", got)
	}
	// no non-5xx declared → nil.
	only5xx := apiRoute{Op: &openapi3.Operation{OperationID: "X"}, Responses: map[string]bool{"500": true}}
	if got := xoh12CheckRoute(only5xx, covered); got != nil {
		t.Errorf("only-5xx should return nil, got %v", got)
	}
}

func TestXoh13CheckGuard_ZeroCov(t *testing.T) {
	fn := ssac.ServiceFunc{Name: "GetX", FileName: "x.ssac"}

	// non-guard type → nil.
	if got := xoh13CheckGuard(fn, ssac.Sequence{Type: "get"}, nil); got != nil {
		t.Error("non-guard should return nil")
	}
	// eval with status 0 → nil (no default).
	if got := xoh13CheckGuard(fn, ssac.Sequence{Type: "eval"}, nil); got != nil {
		t.Error("eval default 0 should return nil")
	}
	// empty guard, default 404, uncovered → warning.
	seq := ssac.Sequence{Type: "empty", Target: "user", Message: "not found", Line: 2}
	diags := xoh13CheckGuard(fn, seq, map[string]bool{})
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d: %+v", len(diags), diags)
	}
	// covered → nil.
	if got := xoh13CheckGuard(fn, seq, map[string]bool{"404": true}); got != nil {
		t.Errorf("covered guard should return nil, got %v", got)
	}
	// explicit ErrStatus honored.
	seq2 := ssac.Sequence{Type: "auth", Action: "delete", Resource: "project", ErrStatus: 403, Line: 3}
	if got := xoh13CheckGuard(fn, seq2, map[string]bool{}); len(got) != 1 {
		t.Errorf("auth guard expected 1 diag, got %v", got)
	}
}

func TestXoh13CheckHappyPath_ZeroCov(t *testing.T) {
	fn := ssac.ServiceFunc{Name: "GetX", FileName: "x.ssac"}

	// no @response sequence → nil.
	if got := xoh13CheckHappyPath(fn, map[string]bool{}); got != nil {
		t.Error("no response seq should return nil")
	}
	// has @response, no 2xx covered → warning.
	fnR := ssac.ServiceFunc{Name: "GetX", FileName: "x.ssac", Sequences: []ssac.Sequence{{Type: "response", Line: 4}}}
	if got := xoh13CheckHappyPath(fnR, map[string]bool{"401": true}); len(got) != 1 {
		t.Fatalf("expected 1 diag, got %v", got)
	}
	// has @response, 2xx covered → nil.
	if got := xoh13CheckHappyPath(fnR, map[string]bool{"200": true}); got != nil {
		t.Errorf("2xx covered should return nil, got %v", got)
	}
}

func TestXoh13CheckFunc_ZeroCov(t *testing.T) {
	// fn with a guard seq and a response seq, no coverage → 2 diags.
	fn := ssac.ServiceFunc{
		Name:     "GetX",
		FileName: "x.ssac",
		Sequences: []ssac.Sequence{
			{Type: "empty", Target: "user", Message: "nf", Line: 2},
			{Type: "response", Line: 5},
		},
	}
	diags := xoh13CheckFunc(fn, map[string]bool{})
	if len(diags) != 2 {
		t.Fatalf("expected 2 diags (guard + happy path), got %d: %+v", len(diags), diags)
	}
}
