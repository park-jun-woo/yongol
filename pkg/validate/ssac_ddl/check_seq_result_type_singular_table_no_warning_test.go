//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what XDS-12 — @result↔DDL 매칭이 단수 정규화(단/복수 명명 모두 허용)로 동작하는지 검증
package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestCheckSeqResultType_SingularTableNoWarning(t *testing.T) {
	fs := &yongol.Fullstack{}
	fn := ssac.ServiceFunc{Name: "GetDashboard", FileName: "GetDashboard.ssac"}
	seq := ssac.Sequence{
		Type:   "get",
		Line:   8,
		Result: &ssac.Result{Type: "AppConfig", Var: "cfg"},
	}
	tables := map[string]bool{canonicalTableKey("app_config"): true}

	diags := checkSeqResultType(fs, tables, fn, seq)
	if len(diags) != 0 {
		t.Fatalf("expected no diagnostics for singular table match, got %d: %+v", len(diags), diags)
	}
}
