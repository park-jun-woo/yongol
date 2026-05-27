//ff:func feature=gen-fastapi type=util control=sequence
//ff:what pySourceExpr — Go dot-access 표현식 → Python dict key access 변환

package ssac

import (
	"fmt"
	"strings"
)

// pySourceExpr converts a Go-style source expression to a Python dict key
// access expression. "token.AccessToken" → "token[\"access_token\"]",
// "workflow" → "workflow".
func pySourceExpr(source string) string {
	dotIdx := strings.Index(source, ".")
	if dotIdx < 0 {
		return source
	}
	obj := source[:dotIdx]
	field := source[dotIdx+1:]
	return fmt.Sprintf("%s[\"%s\"]", obj, snakeCase(field))
}
