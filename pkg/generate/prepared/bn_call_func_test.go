//ff:func feature=generate type=test control=sequence
//ff:what prepared 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package prepared

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func bnCallFunc(prefix string) ssac.ServiceFunc {
	return ssac.ServiceFunc{
		Name:      "f",
		Sequences: []ssac.Sequence{{Type: "call", Model: prefix + "Put"}},
	}
}
