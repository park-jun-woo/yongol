//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvertersZeroCov — 각 SSaC 시퀀스 종류를 BuildServicePlan 으로 변환해 convert* 디스패처 전체 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestConvertState_ZeroCov(t *testing.T) {
	plan := buildPlanOrFail(t, "ArchiveWorkflowState", []ssac.Sequence{
		{Type: ssac.SeqState, DiagramID: "Workflow",
			Inputs:     map[string]string{"Status": "wf.Status"},
			Transition: "ArchiveWorkflow", Message: "Cannot archive", ErrStatus: 409},
	})
	if plan.Ops[0].Kind != OpState {
		t.Errorf("expected OpState, got %d", plan.Ops[0].Kind)
	}
}
