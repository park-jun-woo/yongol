//ff:func feature=gen-fastapi type=util control=selection
//ff:what renderArgValueLegacy — Location 미설정 시 SSaC source 이름 기반 Python 표현식 매핑

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderArgValueLegacy maps SSaC source names when Location is empty.
func renderArgValueLegacy(a ir.FieldArg, col string) string {
	source := a.Source
	if source == "" {
		source = "request"
	}
	if col == "" {
		return legacyColumnlessExpr(source)
	}
	switch source {
	case "request":
		return col
	case "currentUser":
		return fmt.Sprintf("current_user[\"%s\"]", col)
	default:
		return fmt.Sprintf("%s.%s", source, col)
	}
}
