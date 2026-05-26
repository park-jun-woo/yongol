//ff:func feature=gen-nestjs type=util control=sequence
//ff:what renderArgValue — FieldArg → TypeScript 식별자/리터럴 표현식 생성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderArgValue produces a TypeScript expression for a FieldArg.
func renderArgValue(a ir.FieldArg) string {
	if a.Literal != "" {
		if a.IsQuoted {
			return fmt.Sprintf("'%s'", a.Literal)
		}
		return a.Literal
	}
	source := a.Source
	if source == "" {
		source = "params"
	}
	field := strings.TrimPrefix(a.Field, ".")
	if field == "" {
		return source
	}
	return fmt.Sprintf("%s.%s", source, lcFirst(field))
}
