//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what XDS-12 sqlc row type 매칭 회귀 — RowType 등록 시 WARNING 없음, 미등록 시 WARNING

package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
	sqlcparser "github.com/park-jun-woo/yongol/pkg/parser/sqlc"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// TestCheckSeqResultType_SqlcRowTypePasses — JOIN 쿼리의 sqlc 합성 row type 을
// `@get` result 로 사용하면 (DTO 선언 없이, DDL 테이블 매칭 없이) 통과해야 한다.
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

// TestCheckSeqResultType_UnknownTypeStillWarns — 회귀 가드: RowType 에 등록되지
// 않은 합성 타입은 기존대로 WARNING 을 내야 한다 (modelToTable 분기 유지 확인).
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
