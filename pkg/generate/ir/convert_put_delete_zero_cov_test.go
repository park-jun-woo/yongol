//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvertersZeroCov — 각 SSaC 시퀀스 종류를 BuildServicePlan 으로 변환해 convert* 디스패처 전체 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestConvertPutDelete_ZeroCov(t *testing.T) {
	plan := buildPlanOrFail(t, "ArchiveWorkflow", []ssac.Sequence{
		{Type: ssac.SeqPut, Model: "Workflow.UpdateStatus",
			Inputs: map[string]string{"ID": "request.id", "Status": `"archived"`}},
		{Type: ssac.SeqDelete, Model: "Workflow.Delete",
			Inputs: map[string]string{"ID": "request.id"}},
	})
	if plan.Ops[0].Kind != OpPut || plan.Ops[1].Kind != OpDelete {
		t.Errorf("expected OpPut/OpDelete, got %d/%d", plan.Ops[0].Kind, plan.Ops[1].Kind)
	}
}
