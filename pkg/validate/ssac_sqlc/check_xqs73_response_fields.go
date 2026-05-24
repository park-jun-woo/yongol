//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-sqlc
//ff:what checkXqs73ResponseFields — @response 필드 맵의 각 var.Field 참조가 SELECT 컬럼에 존재하는지 검사

package ssac_sqlc

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// checkXqs73ResponseFields checks @response { field: var.Field } field
// references against the partial SELECT column list.
func checkXqs73ResponseFields(fn ssacparser.ServiceFunc, seq ssacparser.Sequence, vars map[string]xqs73VarInfo) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, val := range seq.Fields {
		parts := strings.SplitN(val, ".", 2)
		if len(parts) != 2 {
			continue
		}
		varName, fieldName := parts[0], parts[1]
		vi, ok := vars[varName]
		if !ok {
			continue
		}
		snakeField := toSnake(fieldName)
		if !selectColsContain(vi.query.SelectCols, snakeField) {
			diags = append(diags, xqs73FieldMissingDiag(fn, seq, fieldName, varName, vi))
		}
	}
	return diags
}
