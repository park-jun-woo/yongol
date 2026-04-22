//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @post 파싱 검증 — SeqPost 타입, Result, Inputs 확인

package ssac

import "testing"

func TestParsePost(t *testing.T) {
	src := `package service

// @post Session session = Session.Create({ProjectID: request.ProjectID, Command: request.Command})
func CreateSession(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqPost)
	assertEqual(t, "Result.Type", seq.Result.Type, "Session")
	assertEqual(t, "Result.Var", seq.Result.Var, "session")
	if len(seq.Inputs) != 2 {
		t.Fatalf("expected 2 inputs, got %d", len(seq.Inputs))
	}
}
