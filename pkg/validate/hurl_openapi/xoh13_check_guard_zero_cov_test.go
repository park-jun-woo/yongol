//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what zz_zerocov_test — hurl_openapi 0% 함수 (collectCoveredStatusCodes / xoh12CheckRoute / xoh13CheckFunc / xoh13CheckGuard / xoh13CheckHappyPath) 단위 테스트
package hurl_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

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
