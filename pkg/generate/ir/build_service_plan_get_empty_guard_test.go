//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildServicePlanGetEmptyGuard -- @get + @empty 가드 패턴 → FollowedByGuard 주입 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildServicePlanGetEmptyGuard(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "GetAuditLog",
		FileName: "get_audit_log.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqGet,
				Model: "AuditLog.GetAuditLog",
				Inputs: map[string]string{
					"ID": "request.id",
				},
				Result: &ssac.Result{Var: "detail", Type: "AuditLogGetAuditLogRow"},
			},
			{
				Type:      ssac.SeqEmpty,
				Target:    "detail",
				Message:   "Audit log not found",
				ErrStatus: 404,
			},
			{
				Type:   ssac.SeqResponse,
				Target: "detail",
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan returned error: %v", err)
	}

	if len(plan.Ops) != 3 {
		t.Fatalf("len(Ops) = %d, want 3", len(plan.Ops))
	}

	getOp := plan.Ops[0]
	if getOp.Kind != OpGet {
		t.Fatalf("Ops[0].Kind = %d, want OpGet", getOp.Kind)
	}
	if getOp.Get.FollowedByGuard != OpEmpty {
		t.Errorf("Get.FollowedByGuard = %d, want OpEmpty(%d)", getOp.Get.FollowedByGuard, OpEmpty)
	}

	emptyOp := plan.Ops[1]
	if emptyOp.Kind != OpEmpty {
		t.Fatalf("Ops[1].Kind = %d, want OpEmpty", emptyOp.Kind)
	}
	if emptyOp.Empty.VarName != "detail" {
		t.Errorf("Empty.VarName = %q, want %q", emptyOp.Empty.VarName, "detail")
	}
	if emptyOp.Empty.StatusCode != 404 {
		t.Errorf("Empty.StatusCode = %d, want 404", emptyOp.Empty.StatusCode)
	}
}
