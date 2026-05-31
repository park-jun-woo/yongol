//ff:func feature=validate type=test control=sequence topic=ssac-ddl
//ff:what XDS-12 — @result↔DDL 매칭이 단수 정규화(단/복수 명명 모두 허용)로 동작하는지 검증
package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
