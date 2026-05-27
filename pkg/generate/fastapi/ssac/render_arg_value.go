//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderArgValue — FieldArg → Python 식별자/리터럴 표현식 생성 (source 매핑 포함)

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderArgValue produces a Python expression for a FieldArg.
// SSaC source references are mapped to Python parameter names:
//   - request.{field} → params["{snake_field}"]
//   - request.Params.{field} → params["{snake_field}"] (query params merged)
//   - currentUser.{field} → user["{snake_field}"]
//   - {varName}.{field} → {varName}["{snake_field}"] (result dict access)
func renderArgValue(a ir.FieldArg) string {
	if a.Literal != "" {
		if a.IsQuoted {
			return fmt.Sprintf("\"%s\"", a.Literal)
		}
		return a.Literal
	}
	source := a.Source
	if source == "" {
		source = "params"
	}
	field := strings.TrimPrefix(a.Field, ".")
	if field == "" {
		return mapSource(source)
	}
	return renderSourceField(source, field)
}
