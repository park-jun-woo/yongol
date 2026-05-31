//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvertersZeroCov — 각 SSaC 시퀀스 종류를 BuildServicePlan 으로 변환해 convert* 디스패처 전체 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestConvertAuth_ZeroCov(t *testing.T) {
	plan := buildPlanOrFail(t, "ArchiveWorkflowAuth", []ssac.Sequence{
		{Type: ssac.SeqAuth, Action: "ArchiveWorkflow", Resource: "workflow",
			Inputs: map[string]string{"ResourceID": "wf.ID"}, Message: "Forbidden", ErrStatus: 403},
	})
	if plan.Ops[0].Kind != OpAuth {
		t.Errorf("expected OpAuth, got %d", plan.Ops[0].Kind)
	}
}
