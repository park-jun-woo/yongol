//ff:func feature=generate type=test control=sequence
//ff:what prepared 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package prepared

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestSsacUsesFileCalls_ZeroCov(t *testing.T) {
	if ssacUsesFileCalls(nil) {
		t.Error("nil false")
	}
	if !ssacUsesFileCalls(bnFS(nil, []ssac.ServiceFunc{bnCallFunc("file.")})) {
		t.Error("expected file use")
	}
}
