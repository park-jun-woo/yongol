//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestCollectPlanStubs — collectPlanStubs same-feature 미정의 함수 stub 수집·정의됨/타feature 스킵 검증
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectPlanStubs(t *testing.T) {
	t.Run("CollectsUndefinedSameFeature", func(t *testing.T) {
		plan := &ir.ServicePlan{
			Ops: []ir.Op{
				{Kind: ir.OpCall, Call: &ir.CallOp{Package: "auth", Function: "IssueToken"}},
				{Kind: ir.OpEval, Eval: &ir.EvalOp{Package: "auth", Function: "IsExpired"}},
			},
		}
		stubs := map[string]bool{}
		collectPlanStubs(plan, "auth", map[string]bool{}, stubs)
		if !stubs["issue_token"] || !stubs["is_expired"] {
			t.Errorf("expected issue_token+is_expired stubs, got: %v", stubs)
		}
		if len(stubs) != 2 {
			t.Errorf("expected 2 stubs, got: %v", stubs)
		}
	})

	t.Run("SkipsCrossFeature", func(t *testing.T) {
		plan := &ir.ServicePlan{
			Ops: []ir.Op{
				{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing", Function: "HoldEscrow"}},
			},
		}
		stubs := map[string]bool{}
		collectPlanStubs(plan, "auth", map[string]bool{}, stubs)
		if len(stubs) != 0 {
			t.Errorf("expected no stubs, got: %v", stubs)
		}
	})

	t.Run("SkipsDefined", func(t *testing.T) {
		plan := &ir.ServicePlan{
			Ops: []ir.Op{
				{Kind: ir.OpCall, Call: &ir.CallOp{Package: "auth", Function: "IssueToken"}},
			},
		}
		stubs := map[string]bool{}
		collectPlanStubs(plan, "auth", map[string]bool{"issue_token": true}, stubs)
		if len(stubs) != 0 {
			t.Errorf("expected no stubs for defined func, got: %v", stubs)
		}
	})

	t.Run("SkipsEmptyTarget", func(t *testing.T) {
		plan := &ir.ServicePlan{
			Ops: []ir.Op{{Kind: ir.OpAuth}},
		}
		stubs := map[string]bool{}
		collectPlanStubs(plan, "auth", map[string]bool{}, stubs)
		if len(stubs) != 0 {
			t.Errorf("expected no stubs for non-call op, got: %v", stubs)
		}
	})
}
