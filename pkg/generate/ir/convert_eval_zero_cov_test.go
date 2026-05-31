//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvertersZeroCov — 각 SSaC 시퀀스 종류를 BuildServicePlan 으로 변환해 convert* 디스패처 전체 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestConvertEval_ZeroCov(t *testing.T) {
	plan := buildPlanOrFail(t, "CheckSchedule", []ssac.Sequence{
		{Type: ssac.SeqEval, Model: "schedule.IsExpired",
			Inputs:  map[string]string{"StartsAt": "request.starts_at"},
			Message: "Schedule date is in the past", ErrStatus: 400},
	})
	if plan.Ops[0].Kind != OpEval {
		t.Errorf("expected OpEval, got %d", plan.Ops[0].Kind)
	}
}
