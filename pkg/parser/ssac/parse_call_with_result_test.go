//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @call 결과 있음 파싱 검증 — Result.Type, Result.Var, Inputs 확인

package ssac

import "testing"

func TestParseCallWithResult(t *testing.T) {
	src := `package service

// @call Token token = auth.VerifyPassword({Email: user.Email, Password: request.Password})
func Login(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqCall)
	assertEqual(t, "Model", seq.Model, "auth.VerifyPassword")
	if seq.Result == nil {
		t.Fatal("expected result")
	}
	assertEqual(t, "Result.Type", seq.Result.Type, "Token")
	assertEqual(t, "Result.Var", seq.Result.Var, "token")
	if len(seq.Inputs) != 2 {
		t.Fatalf("expected 2 inputs, got %d", len(seq.Inputs))
	}
	assertEqual(t, "Inputs.Email", seq.Inputs["Email"], "user.Email")
	assertEqual(t, "Inputs.Password", seq.Inputs["Password"], "request.Password")
}
