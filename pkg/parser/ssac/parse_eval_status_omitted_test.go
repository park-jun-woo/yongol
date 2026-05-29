//ff:func feature=ssac-parse type=test control=sequence
//ff:what TestParseEvalStatusOmittedCapturedAsZero — @eval STATUS 누락 시 파서는 통과(0); S-68 이 ERROR

package ssac

import "testing"

func TestParseEvalStatusOmittedCapturedAsZero(t *testing.T) {
	src := `package service

// @eval pkg.IsThing({}) "msg"
func Anything(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqEval)
	assertEqual(t, "Message", seq.Message, "msg")
	if seq.ErrStatus != 0 {
		t.Errorf("expected ErrStatus 0 (omitted) before S-68 validation, got %d", seq.ErrStatus)
	}
}
