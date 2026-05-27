//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestCollectSameFeatureStubs — 같은 feature @call/@eval 대상 inline stub 수집 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectSameFeatureStubs(t *testing.T) {
	t.Run("SameFeatureCallCollected", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "Login",
				Feature:     "auth",
				Ops: []ir.Op{
					{Kind: ir.OpCall, Call: &ir.CallOp{Package: "auth", Function: "IssueToken"}},
				},
			},
		}
		got := collectSameFeatureStubs(plans, "auth")
		if len(got) != 1 || got[0] != "issue_token" {
			t.Errorf("expected [issue_token], got: %v", got)
		}
	})

	t.Run("CrossFeatureCallSkipped", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "CreateOrder",
				Feature:     "order",
				Ops: []ir.Op{
					{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing", Function: "HoldEscrow"}},
				},
			},
		}
		got := collectSameFeatureStubs(plans, "order")
		if len(got) != 0 {
			t.Errorf("expected no stubs for cross-feature call, got: %v", got)
		}
	})

	t.Run("DefinedFunctionNotStubbed", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "IssueToken",
				Feature:     "auth",
				Ops:         []ir.Op{},
			},
			{
				OperationID: "Login",
				Feature:     "auth",
				Ops: []ir.Op{
					{Kind: ir.OpCall, Call: &ir.CallOp{Package: "auth", Function: "IssueToken"}},
				},
			},
		}
		got := collectSameFeatureStubs(plans, "auth")
		if len(got) != 0 {
			t.Errorf("expected no stub for defined function, got: %v", got)
		}
	})

	t.Run("SameFeatureEvalCollected", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "RunSchedule",
				Feature:     "schedule",
				Ops: []ir.Op{
					{Kind: ir.OpEval, Eval: &ir.EvalOp{Package: "schedule", Function: "ParseCron"}},
				},
			},
		}
		got := collectSameFeatureStubs(plans, "schedule")
		if len(got) != 1 || got[0] != "parse_cron" {
			t.Errorf("expected [parse_cron], got: %v", got)
		}
	})

	t.Run("MultipleStubsSorted", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "ViewDashboard",
				Feature:     "dashboard",
				Ops: []ir.Op{
					{Kind: ir.OpCall, Call: &ir.CallOp{Package: "dashboard", Function: "Summarize"}},
					{Kind: ir.OpCall, Call: &ir.CallOp{Package: "dashboard", Function: "BuildExecutionDetail"}},
				},
			},
		}
		got := collectSameFeatureStubs(plans, "dashboard")
		if len(got) != 2 {
			t.Fatalf("expected 2 stubs, got: %v", got)
		}
		if got[0] != "build_execution_detail" || got[1] != "summarize" {
			t.Errorf("expected sorted [build_execution_detail, summarize], got: %v", got)
		}
	})

	t.Run("EmptyPlans", func(t *testing.T) {
		got := collectSameFeatureStubs(nil, "auth")
		if len(got) != 0 {
			t.Errorf("expected no stubs for nil plans, got: %v", got)
		}
	})
}
