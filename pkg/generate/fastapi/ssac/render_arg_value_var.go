//ff:func feature=gen-fastapi type=util control=sequence
//ff:what renderArgValueVar — LocVar FieldArg 를 source.col Python 접근 표현식으로 렌더

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderArgValueVar renders a LocVar FieldArg. SourceColumn is the snake_case
// name on the source variable/struct; it falls back to the field name and then
// the target ColumnName.
func renderArgValueVar(a ir.FieldArg, col string) string {
	srcCol := a.SourceColumn
	if srcCol == "" {
		srcCol = snakeCase(fieldName(a))
	}
	if srcCol == "" {
		srcCol = col
	}
	if srcCol != "" && a.Source != "" {
		return fmt.Sprintf("%s.%s", a.Source, srcCol)
	}
	if a.Source != "" {
		return a.Source
	}
	return srcCol
}
