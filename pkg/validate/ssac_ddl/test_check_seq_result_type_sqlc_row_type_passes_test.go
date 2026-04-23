//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what XDS-12 — 등록된 sqlc RowType 는 @get 결과로 사용해도 진단 없음

package ssac_ddl

import (
	"testing"

	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
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
