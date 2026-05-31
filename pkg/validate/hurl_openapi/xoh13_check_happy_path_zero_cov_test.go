//ff:func feature=validate type=test control=sequence topic=hurl-openapi
//ff:what zz_zerocov_test — hurl_openapi 0% 함수 (collectCoveredStatusCodes / xoh12CheckRoute / xoh13CheckFunc / xoh13CheckGuard / xoh13CheckHappyPath) 단위 테스트
package hurl_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

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
