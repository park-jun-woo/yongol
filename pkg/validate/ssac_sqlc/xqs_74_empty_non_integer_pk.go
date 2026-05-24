//ff:func feature=validate type=rule control=iteration dimension=1 topic=ssac-sqlc
//ff:what XQS-74 — @empty/@exists guard 대상 모델의 PK가 integer가 아니면 ERROR

package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xqs74IntegerTypes lists PG types whose Go zero value is 0 (integer family).
var xqs74IntegerTypes = map[string]bool{
	"BIGINT":    true,
	"INT":       true,
	"INTEGER":   true,
	"SERIAL":    true,
	"BIGSERIAL": true,
	"INT2":      true,
	"INT4":      true,
	"INT8":      true,
	"SMALLINT":  true,
}

// xqs74EmptyNonIntegerPK validates XQS-74: @empty/@exists guards whose target
// variable was bound by a preceding @get/@post sequence must reference a model
// whose DDL PK column is an integer type.
func xqs74EmptyNonIntegerPK(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || len(fs.ServiceFuncs) == 0 || len(fs.DDLTables) == 0 {
		return nil
	}
	tableMap := buildDDLTableLookup(fs)
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		diags = append(diags, xqs74CheckFunc(fn, tableMap)...)
	}
	return diags
}
