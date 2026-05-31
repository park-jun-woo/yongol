//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvertersZeroCov — 각 SSaC 시퀀스 종류를 BuildServicePlan 으로 변환해 convert* 디스패처 전체 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestConvertGet_ZeroCov(t *testing.T) {
	plan := buildPlanOrFail(t, "GetCourse", []ssac.Sequence{
		{Type: ssac.SeqGet, Model: "Course.FindByID",
			Inputs: map[string]string{"ID": "request.id"},
			Result: &ssac.Result{Var: "course", Type: "Course"}},
		{Type: ssac.SeqResponse, Target: "course"},
	})
	if plan.Ops[0].Kind != OpGet {
		t.Errorf("expected OpGet, got %d", plan.Ops[0].Kind)
	}
}
