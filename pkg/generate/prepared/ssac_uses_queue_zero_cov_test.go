//ff:func feature=generate type=test control=sequence
//ff:what prepared 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package prepared

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestSsacUsesQueue_ZeroCov(t *testing.T) {
	if ssacUsesQueue(nil) {
		t.Error("nil false")
	}
	pub := ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "publish"}}}
	if !ssacUsesQueue(bnFS(nil, []ssac.ServiceFunc{pub})) {
		t.Error("expected queue use")
	}
}
