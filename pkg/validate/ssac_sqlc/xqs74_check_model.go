//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what xqs74CheckModel — 모델의 DDL PK 컬럼 타입이 integer인지 검사하여 진단 반환

package ssac_sqlc

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xqs74CheckModel looks up the DDL table for the model and checks if its PK
// column is an integer type. Returns (diag, true) on violation.
func xqs74CheckModel(seq ssacparser.Sequence, modelName string, tableMap map[string]*ddl.Table, fileName string) (diagnostic.Diagnostic, bool) {
	tableName := modelToTableName(modelName)
	table, ok := tableMap[tableName]
	if !ok {
		return diagnostic.Diagnostic{}, false
	}
	if len(table.PrimaryKey) == 0 {
		return diagnostic.Diagnostic{}, false
	}
	pkColName := table.PrimaryKey[0]
	col, ok := table.Columns[pkColName]
	if !ok {
		return diagnostic.Diagnostic{}, false
	}
	rawUpper := strings.ToUpper(col.RawType)
	if xqs74IntegerTypes[rawUpper] {
		return diagnostic.Diagnostic{}, false
	}
	return diagnostic.Diagnostic{
		File:  fileName,
		Line:  seq.Line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XQS-74] @%s guard on %s requires integer PK but table %s has %s PK (%s)",
			seq.Type, modelName, tableName, col.RawType, pkColName,
		),
		Advice: fmt.Sprintf(
			"Add id BIGINT PRIMARY KEY column to %s",
			tableName,
		),
	}, true
}
