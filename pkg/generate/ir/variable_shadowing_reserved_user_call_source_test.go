//ff:func feature=gen-ir type=test control=sequence
//ff:what TestVariableShadowing -- 동일 VarName 중복 선언 시 _result 접미사 자동 해소 + 후속 Op Source 갱신 검증
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestVariableShadowingReservedUserCallSource(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "Login",
		FileName: "login.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:         ssac.SeqVerifyPassword,
				Model:        "User",
				EmailCol:     "Email",
				EmailExpr:    "request.Email",
				HashCol:      "PasswordHash",
				PasswordExpr: "request.Password",
				Result:       &ssac.Result{Var: "user", Type: "User"},
			},
			{
				Type:   ssac.SeqCall,
				Model:  "auth.IssueToken",
				Inputs: map[string]string{"Email": "user.Email", "Role": "user.Role"},
				Result: &ssac.Result{Var: "token", Type: "TokenPair"},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	// VerifyPassword result "user" should be renamed to "user_result"
	// because "user" is a reserved name.
	vpOp := plan.Ops[0]
	if vpOp.Kind != OpVerifyPassword {
		t.Fatalf("Ops[0].Kind = %d, want OpVerifyPassword", vpOp.Kind)
	}
	if vpOp.VerifyPW.ResultVar != "user_result" {
		t.Errorf("VerifyPW.ResultVar = %q, want %q", vpOp.VerifyPW.ResultVar, "user_result")
	}

	// @call args should have Source updated from "user" to "user_result".
	callOp := plan.Ops[1]
	if callOp.Kind != OpCall {
		t.Fatalf("Ops[1].Kind = %d, want OpCall", callOp.Kind)
	}
	emailArg := findArgByKey(callOp.Call.Args, "Email")
	if emailArg == nil {
		t.Fatal("missing Email arg")
	}
	if emailArg.Source != "user_result" {
		t.Errorf("Email.Source = %q, want %q", emailArg.Source, "user_result")
	}
	roleArg := findArgByKey(callOp.Call.Args, "Role")
	if roleArg == nil {
		t.Fatal("missing Role arg")
	}
	if roleArg.Source != "user_result" {
		t.Errorf("Role.Source = %q, want %q", roleArg.Source, "user_result")
	}
}
