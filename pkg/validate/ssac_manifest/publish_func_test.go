//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what XSA-71/72/74-77 + Run test — backend required(ERROR)/unused(WARNING) 규칙 검증
package ssac_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func publishFunc() ssac.ServiceFunc {
	return ssac.ServiceFunc{Sequences: []ssac.Sequence{{Type: "publish"}}}
}
