//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what xqs73FieldMissingDiag — 부분 SELECT에서 필드 누락 시 진단 메시지 생성

package ssac_sqlc

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xqs73FieldMissingDiag creates a diagnostic for a field reference that is
// missing from the partial SELECT column list.
func xqs73FieldMissingDiag(fn ssacparser.ServiceFunc, seq ssacparser.Sequence, fieldName, varName string, vi xqs73VarInfo) diagnostic.Diagnostic {
	snakeField := toSnake(fieldName)
	return diagnostic.Diagnostic{
		File:    fn.FileName,
		Line:    seq.Line,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: fmt.Sprintf("[XQS-73] field %q referenced on query result %q, but query %q does not SELECT this column (partial SELECT: %s)", fieldName, varName, vi.query.Name, strings.Join(vi.query.SelectCols, ", ")),
		Advice:  fmt.Sprintf("Add %q to the SELECT column list of query %q, or use SELECT *.", snakeField, vi.query.Name),
	}
}
