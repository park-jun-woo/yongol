//ff:func feature=ssac-parse type=test control=sequence
//ff:what TestParseEvalBasic — @eval 단일 인자 + message + status 파싱

package ssac

import "testing"

func TestParseEvalBasic(t *testing.T) {
	src := `package service

// @eval billing.IsZeroBalance({Balance: org.CreditsBalance}) "Insufficient credits" 402
func ChargeOrder(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqEval)
	assertEqual(t, "Model", seq.Model, "billing.IsZeroBalance")
	assertEqual(t, "Message", seq.Message, "Insufficient credits")
	if seq.ErrStatus != 402 {
		t.Errorf("expected ErrStatus 402, got %d", seq.ErrStatus)
	}
	if seq.Result != nil {
		t.Error("expected no result for @eval")
	}
	if len(seq.Inputs) != 1 || seq.Inputs["Balance"] != "org.CreditsBalance" {
		t.Errorf("expected single Balance input, got %v", seq.Inputs)
	}
}
