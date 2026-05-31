//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what XDS-12 — @result↔DDL 매칭이 단수 정규화(단/복수 명명 모두 허용)로 동작하는지 검증
package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCheckSeqResultType_PluralTableNoWarning(t *testing.T) {
	fs := &yongol.Fullstack{}
	fn := ssac.ServiceFunc{Name: "GetUser", FileName: "GetUser.ssac"}
	seq := ssac.Sequence{
		Type:   "get",
		Line:   3,
		Result: &ssac.Result{Type: "User", Var: "u"},
	}
	tables := map[string]bool{canonicalTableKey("users"): true}

	diags := checkSeqResultType(fs, tables, fn, seq)
	if len(diags) != 0 {
		t.Fatalf("expected no diagnostics for plural table match, got %d: %+v", len(diags), diags)
	}
}
