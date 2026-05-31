//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderLiteralArg — FieldArg 리터럴 값을 TypeScript 표현식으로 렌더 (quoted 여부 반영)

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderLiteralArg renders a FieldArg literal as a TypeScript expression,
// quoting it when the literal was a quoted string.
func renderLiteralArg(a ir.FieldArg) string {
	if a.IsQuoted {
		return fmt.Sprintf("'%s'", a.Literal)
	}
	return a.Literal
}
