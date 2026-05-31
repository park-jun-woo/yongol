//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestVariableShadowing -- 동일 VarName 중복 선언 시 _result 접미사 자동 해소 + 후속 Op Source 갱신 검증
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestVariableShadowingResponseSourceUpdate(t *testing.T) {
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
				Inputs: map[string]string{"Email": "user.Email"},
				Result: &ssac.Result{Var: "token", Type: "TokenPair"},
			},
			{
				Type: ssac.SeqResponse,
				Fields: map[string]string{
					"access_token":  "token.AccessToken",
					"refresh_token": "token.RefreshToken",
					"email":         "user.Email",
				},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	// Response op should have "user" references updated to "user_result".
	respOp := plan.Ops[2]
	if respOp.Kind != OpResponse {
		t.Fatalf("Ops[2].Kind = %d, want OpResponse", respOp.Kind)
	}

	for _, f := range respOp.Response.Fields {
		assertResponseFieldSource(t, f.Name, f.Source)
	}
}
