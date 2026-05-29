//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what XDS-12 — 등록되지 않은 복합 타입은 여전히 WARNING

package ssac_ddl

import (
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestCheckSeqResultType_UnknownTypeStillWarns — regression guard: a composite type not
// registered in RowType must still emit a WARNING (verifies the modelToTable branch is preserved).
func TestCheckSeqResultType_UnknownTypeStillWarns(t *testing.T) {
	fs := &yongol.Fullstack{
		SQLcQueries: []sqlcparser.QuerySpec{
			{Name: "GetAuditLogDetail", Cardinality: "one", RowType: "GetAuditLogDetailRow"},
		},
	}
	fn := ssac.ServiceFunc{Name: "GetSomething", FileName: "x.ssac"}
	seq := ssac.Sequence{
		Type: "get",
		Line: 20,
		Result: &ssac.Result{
			Type: "NotARealRow",
			Var:  "row",
		},
	}
	tables := map[string]bool{}

	diags := checkSeqResultType(fs, tables, fn, seq)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic for unknown type, got %d: %+v", len(diags), diags)
	}
}
