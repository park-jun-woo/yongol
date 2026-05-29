//ff:func feature=ssac-parse type=test control=sequence
//ff:what TestParseEvalMultiInput — @eval 다중 인자 (string literal 포함) 파싱

package ssac

import "testing"

func TestParseEvalMultiInput(t *testing.T) {
	src := `package service

// @eval rate.IsLimited({UserID: currentUser.ID, Endpoint: "ExecuteWorkflow"}) "Rate limited" 429
func RunWorkflow(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqEval)
	assertEqual(t, "Model", seq.Model, "rate.IsLimited")
	if len(seq.Inputs) != 2 {
		t.Fatalf("expected 2 inputs, got %d", len(seq.Inputs))
	}
	assertEqual(t, "Inputs.UserID", seq.Inputs["UserID"], "currentUser.ID")
	assertEqual(t, "Inputs.Endpoint", seq.Inputs["Endpoint"], `"ExecuteWorkflow"`)
	assertEqual(t, "Message", seq.Message, "Rate limited")
	if seq.ErrStatus != 429 {
		t.Errorf("expected ErrStatus 429, got %d", seq.ErrStatus)
	}
}
