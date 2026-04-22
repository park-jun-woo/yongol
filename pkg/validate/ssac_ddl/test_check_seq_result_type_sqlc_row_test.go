//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what XDS-12 sqlc row type matching regression — no WARNING when RowType is registered, WARNING when it is not

package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestCheckSeqResultType_SqlcRowTypePasses — using a sqlc composite row type from a JOIN query
// as a `@get` result (with no DTO declaration and no DDL table match) must pass with zero diagnostics.
func TestCheckSeqResultType_SqlcRowTypePasses(t *testing.T) {
	fs := &yongol.Fullstack{
		SQLcQueries: []sqlcparser.QuerySpec{
			{Name: "GetAuditLogDetail", Cardinality: "one", RowType: "GetAuditLogDetailRow"},
		},
	}
	fn := ssac.ServiceFunc{Name: "GetAuditLog", FileName: "audit.ssac"}
	seq := ssac.Sequence{
		Type: "get",
		Line: 10,
		Result: &ssac.Result{
			Type: "GetAuditLogDetailRow",
			Var:  "row",
		},
	}
	tables := map[string]bool{}

	diags := checkSeqResultType(fs, tables, fn, seq)
	if len(diags) != 0 {
		t.Fatalf("expected no diagnostics for sqlc row type, got %d: %+v", len(diags), diags)
	}
}

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
