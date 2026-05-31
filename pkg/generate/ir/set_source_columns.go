//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what setSourceColumns -- FieldArg.Field 접근자에서 SourceColumn(snake_case) 파생

package ir

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/util/caseconv"
)

// setSourceColumns derives SourceColumn from each FieldArg.Field accessor.
func setSourceColumns(args []FieldArg) {
	for j := range args {
		if args[j].Field != "" {
			args[j].SourceColumn = caseconv.PascalToSnake(strings.TrimPrefix(args[j].Field, "."))
		}
	}
}
