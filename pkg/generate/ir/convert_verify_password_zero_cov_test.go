//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvertersZeroCov — 각 SSaC 시퀀스 종류를 BuildServicePlan 으로 변환해 convert* 디스패처 전체 커버
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestConvertVerifyPassword_ZeroCov(t *testing.T) {
	plan := buildPlanOrFail(t, "Login", []ssac.Sequence{
		{Type: ssac.SeqVerifyPassword, Model: "User",
			EmailCol: "email", EmailExpr: "request.body.email",
			HashCol: "password_hash", PasswordExpr: "request.body.password",
			Result:    &ssac.Result{Var: "user", Type: "User"},
			ErrStatus: 401, Message: "Invalid credentials"},
		{Type: ssac.SeqResponse, Target: "user"},
	})
	if plan.Ops[0].Kind != OpVerifyPassword {
		t.Errorf("expected OpVerifyPassword, got %d", plan.Ops[0].Kind)
	}
}
