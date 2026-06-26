//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what zz_zerocov_test — hurl_openapi 0% 함수 (collectCoveredStatusCodes / xoh12CheckRoute / xoh13CheckFunc / xoh13CheckGuard / xoh13CheckHappyPath) 단위 테스트
package hurl_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

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
	diags := xoh13CheckFunc(fn, map[string]bool{}, nil)
	if len(diags) != 2 {
		t.Fatalf("expected 2 diags (guard + happy path), got %d: %+v", len(diags), diags)
	}
}
