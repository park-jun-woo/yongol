//ff:func feature=validate type=test topic=ssac-ddl
//ff:what zz_zerocov_test — ssac_ddl 0% (Run / xds12ResultNoDDLTable / collectFuncResultDiags) 단위 테스트
package ssac_ddl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func fsWithResultZeroCov() *yongol.Fullstack {
	// One @get sequence binding a @result of type "Workflow" with no DDL
	// table → XDS-12 warning.
	return &yongol.Fullstack{
		ServiceFuncs: []ssac.ServiceFunc{{
			Name:     "GetWorkflow",
			FileName: "svc/get.ssac",
			Sequences: []ssac.Sequence{{
				Type:   "get",
				Line:   2,
				Result: &ssac.Result{Type: "Workflow", Var: "wf"},
			}},
		}},
	}
}

func TestRun_ZeroCov(t *testing.T) {
	if diags := Run(&yongol.Fullstack{}); len(diags) != 0 {
		t.Fatalf("empty fullstack → 0 diags, got %d: %+v", len(diags), diags)
	}
}

func TestXds12ResultNoDDLTable_ZeroCov(t *testing.T) {
	diags := xds12ResultNoDDLTable(fsWithResultZeroCov())
	if len(diags) != 1 {
		t.Fatalf("expected 1 XDS-12 diag, got %d: %+v", len(diags), diags)
	}
}

func TestCollectFuncResultDiags_ZeroCov(t *testing.T) {
	fs := fsWithResultZeroCov()
	fn := fs.ServiceFuncs[0]
	tables := canonicalDDLTableSet(fs)
	diags := collectFuncResultDiags(fs, tables, fn)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diag, got %d: %+v", len(diags), diags)
	}

	// Skip branches: call type, package set, nil/empty result.
	skip := ssac.ServiceFunc{Sequences: []ssac.Sequence{
		{Type: "call", Result: &ssac.Result{Type: "Workflow"}},
		{Package: "session", Result: &ssac.Result{Type: "Workflow"}},
		{Type: "get", Result: nil},
		{Type: "get", Result: &ssac.Result{Type: ""}},
	}}
	if got := collectFuncResultDiags(fs, tables, skip); len(got) != 0 {
		t.Errorf("skip branches should yield 0 diags, got %d: %+v", len(got), got)
	}
}
