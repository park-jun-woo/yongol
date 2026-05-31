//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvertersZeroCov — 각 SSaC 시퀀스 종류를 BuildServicePlan 으로 변환해 convert* 디스패처 전체 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestConvertPost_ZeroCov(t *testing.T) {
	plan := buildPlanOrFail(t, "CreateCourse", []ssac.Sequence{
		{Type: ssac.SeqPost, Model: "Course.Create",
			Inputs: map[string]string{"Title": "request.title", "InstructorID": "currentUser.ID"},
			Result: &ssac.Result{Var: "course", Type: "Course"}},
		{Type: ssac.SeqResponse, Fields: map[string]string{"course": "course"}},
	})
	if plan.Ops[0].Kind != OpPost {
		t.Errorf("expected OpPost, got %d", plan.Ops[0].Kind)
	}
}
