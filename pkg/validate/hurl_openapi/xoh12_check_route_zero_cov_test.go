//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what zz_zerocov_test — hurl_openapi 0% 함수 (collectCoveredStatusCodes / xoh12CheckRoute / xoh13CheckFunc / xoh13CheckGuard / xoh13CheckHappyPath) 단위 테스트
package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

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
