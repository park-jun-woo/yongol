//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildServicePlanGetExistsGuard -- @get + @exists 가드 패턴 → FollowedByGuard=OpExists 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildServicePlanGetExistsGuard(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "PublishTemplate",
		FileName: "publish_template.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqGet,
				Model: "Template.FindBySourceWorkflowID",
				Inputs: map[string]string{
					"SourceWorkflowID": "wf.ID",
				},
				Result: &ssac.Result{Var: "existing", Type: "Template"},
			},
			{
				Type:      ssac.SeqExists,
				Target:    "existing",
				Message:   "Already published",
				ErrStatus: 409,
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan returned error: %v", err)
	}

	if len(plan.Ops) != 2 {
		t.Fatalf("len(Ops) = %d, want 2", len(plan.Ops))
	}

	getOp := plan.Ops[0]
	if getOp.Get.FollowedByGuard != OpExists {
		t.Errorf("Get.FollowedByGuard = %d, want OpExists(%d)", getOp.Get.FollowedByGuard, OpExists)
	}

	existsOp := plan.Ops[1]
	if existsOp.Kind != OpExists {
		t.Fatalf("Ops[1].Kind = %d, want OpExists", existsOp.Kind)
	}
	if existsOp.Exists.StatusCode != 409 {
		t.Errorf("Exists.StatusCode = %d, want 409", existsOp.Exists.StatusCode)
	}
}
