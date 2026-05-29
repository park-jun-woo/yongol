//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @empty ErrStatus 파싱 검증 — 커스텀 HTTP 상태 코드 402

package ssac

import "testing"

func TestParseEmptyErrStatus(t *testing.T) {
	src := `package service

// @empty orgWithCredits "Insufficient credits" 402
func ActivateWorkflow() {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqEmpty)
	assertEqual(t, "Target", seq.Target, "orgWithCredits")
	assertEqual(t, "Message", seq.Message, "Insufficient credits")
	if seq.ErrStatus != 402 {
		t.Errorf("expected ErrStatus 402, got %d", seq.ErrStatus)
	}
}
