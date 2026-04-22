//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @state ErrStatus 파싱 검증 — 커스텀 HTTP 상태 코드 422

package ssac

import "testing"

func TestParseStateErrStatus(t *testing.T) {
	src := `package service

// @state workflow {status: workflow.Status} "ActivateWorkflow" "Cannot transition" 422
func ActivateWorkflow() {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqState)
	assertEqual(t, "DiagramID", seq.DiagramID, "workflow")
	assertEqual(t, "Transition", seq.Transition, "ActivateWorkflow")
	assertEqual(t, "Message", seq.Message, "Cannot transition")
	if seq.ErrStatus != 422 {
		t.Errorf("expected ErrStatus 422, got %d", seq.ErrStatus)
	}
}
