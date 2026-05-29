//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what XDS-12 — @result↔DDL 매칭이 단수 정규화(단/복수 명명 모두 허용)로 동작하는지 검증

package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestCheckSeqResultType_SingularTableNoWarning — a singular-named DDL table
// (app_config) matching model AppConfig in singular context must not warn (XSD-55 일관).
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

// TestCheckSeqResultType_PluralTableNoWarning — a plural-named DDL table (users)
// matching model User in singular context must still match (no regression).
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

// TestCheckSeqResultType_PluralInSingularContextStillErrors — Check #1 must keep
// firing for a wrapper-less plural element type, even when a matching table exists.
func TestCheckSeqResultType_PluralInSingularContextStillErrors(t *testing.T) {
	fs := &yongol.Fullstack{}
	fn := ssac.ServiceFunc{Name: "GetWorkflows", FileName: "GetWorkflows.ssac"}
	seq := ssac.Sequence{
		Type:   "get",
		Line:   8,
		Result: &ssac.Result{Type: "Workflows", Var: "wf"},
	}
	tables := map[string]bool{canonicalTableKey("workflows"): true}

	diags := checkSeqResultType(fs, tables, fn, seq)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic for plural-in-singular-context, got %d: %+v", len(diags), diags)
	}
	if diags[0].Level != diagnostic.LevelError {
		t.Fatalf("expected ERROR level (Check #1), got %q", diags[0].Level)
	}
}

// TestCheckSeqResultType_NoCanonicalMatchStillWarns — a type that matches no
// table by canonical form must still emit a WARNING (rule purpose preserved).
func TestCheckSeqResultType_NoCanonicalMatchStillWarns(t *testing.T) {
	fs := &yongol.Fullstack{}
	fn := ssac.ServiceFunc{Name: "GetGhost", FileName: "GetGhost.ssac"}
	seq := ssac.Sequence{
		Type:   "get",
		Line:   5,
		Result: &ssac.Result{Type: "Ghost", Var: "g"},
	}
	tables := map[string]bool{canonicalTableKey("users"): true}

	diags := checkSeqResultType(fs, tables, fn, seq)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic for unmatched type, got %d: %+v", len(diags), diags)
	}
}
