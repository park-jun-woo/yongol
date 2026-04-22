//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what inputValueRefBase — input 값 표현에서 선두 변수 참조만 추출 (".field" 제거)

package ssac

import (
	"strings"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// inputValueRefBase strips a possible ".field" suffix from an input value
// expression and returns the leading variable reference (or "" for literals).
func inputValueRefBase(val string) string {
	if val == "" {
		return ""
	}
	if strings.HasPrefix(val, `"`) || parsessac.IsLiteral(val) {
		return ""
	}
	return strings.SplitN(val, ".", 2)[0]
}
