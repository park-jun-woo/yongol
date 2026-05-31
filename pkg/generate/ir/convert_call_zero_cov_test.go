//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvertersZeroCov — 각 SSaC 시퀀스 종류를 BuildServicePlan 으로 변환해 convert* 디스패처 전체 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestConvertCall_ZeroCov(t *testing.T) {
	plan := buildPlanOrFail(t, "RefreshToken", []ssac.Sequence{
		{Type: ssac.SeqCall, Model: "auth.RefreshRotate",
			Inputs: map[string]string{"RefreshToken": "request.refresh_token"},
			Result: &ssac.Result{Var: "rotated", Type: "auth.RefreshRotateResponse"}},
		{Type: ssac.SeqResponse, Fields: map[string]string{
			"access_token": "rotated.AccessToken", "refresh_token": "rotated.RefreshToken"}},
	})
	if plan.Ops[0].Kind != OpCall {
		t.Errorf("expected OpCall, got %d", plan.Ops[0].Kind)
	}
}
