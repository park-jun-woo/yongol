//ff:func feature=validate type=test control=sequence dimension=1 topic=ssac-structural
//ff:what S-37 scalar skip — @get int64 FK 참조 시 S-37 skip 확인

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestS37ScalarResultSkips(t *testing.T) {
	fs := &yongol.Fullstack{
		ServiceFuncs: []parsessac.ServiceFunc{{
			Name:     "ListExecutionLogs",
			FileName: "service/listexecutionlogs.ssac",
			Sequences: []parsessac.Sequence{
				{
					Type:   "get",
					Model:  "Workflow.FindByID",
					Result: &parsessac.Result{Type: "Workflow", Var: "wf"},
					Args:   []parsessac.Arg{{Field: "ID", Source: "request"}},
					Line:   10,
				},
				{
					Type:    "empty",
					Target:  "wf",
					Message: "Workflow not found",
					Line:    11,
				},
				{
					Type:   "get",
					Model:  "ExecutionLog.Count",
					Result: &parsessac.Result{Type: "int64", Var: "total"},
					Args:   []parsessac.Arg{{Field: "WorkflowID", Source: "wf"}},
					Line:   12,
				},
			},
		}},
	}
	g := &rule.Ground{
		Lookup: map[string]rule.StringSet{},
		Types:  map[string]string{},
		Pairs:  map[string]rule.StringSet{},
		Config: map[string]bool{},
		Vars:   rule.StringSet{},
		Flags:  rule.StringSet{},
	}
	fs.SetGround(g)

	diags := s37FKReferenceGuard(fs)
	if len(diags) != 0 {
		t.Errorf("expected 0 diags for scalar int64 result, got %d: %+v", len(diags), diags)
	}
}
