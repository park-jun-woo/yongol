//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildServicePlanVerifyPassword -- @verify-password 시퀀스 IR 변환 (로그인 타이밍 방어)

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildServicePlanVerifyPassword(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "Login",
		FileName: "login.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:         ssac.SeqVerifyPassword,
				Model:        "User",
				EmailCol:     "email",
				EmailExpr:    "request.body.email",
				HashCol:      "password_hash",
				PasswordExpr: "request.body.password",
				Result:       &ssac.Result{Var: "user", Type: "User"},
				ErrStatus:    401,
				Message:      "Invalid credentials",
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan returned error: %v", err)
	}

	if len(plan.Ops) != 1 {
		t.Fatalf("len(Ops) = %d, want 1", len(plan.Ops))
	}

	vpOp := plan.Ops[0]
	if vpOp.Kind != OpVerifyPassword {
		t.Fatalf("Ops[0].Kind = %d, want OpVerifyPassword", vpOp.Kind)
	}
	if vpOp.VerifyPW.Model != "User" {
		t.Errorf("VerifyPW.Model = %q, want %q", vpOp.VerifyPW.Model, "User")
	}
	if vpOp.VerifyPW.EmailCol != "email" {
		t.Errorf("VerifyPW.EmailCol = %q, want %q", vpOp.VerifyPW.EmailCol, "email")
	}
	if vpOp.VerifyPW.HashCol != "password_hash" {
		t.Errorf("VerifyPW.HashCol = %q, want %q", vpOp.VerifyPW.HashCol, "password_hash")
	}
	if vpOp.VerifyPW.ResultVar != "user" {
		t.Errorf("VerifyPW.ResultVar = %q, want %q", vpOp.VerifyPW.ResultVar, "user")
	}
	if vpOp.VerifyPW.ErrStatus != 401 {
		t.Errorf("VerifyPW.ErrStatus = %d, want 401", vpOp.VerifyPW.ErrStatus)
	}
}
