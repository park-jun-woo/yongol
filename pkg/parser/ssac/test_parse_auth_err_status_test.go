//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @auth ErrStatus 파싱 검증 — 커스텀 HTTP 상태 코드 401

package ssac

import "testing"

func TestParseAuthErrStatus(t *testing.T) {
	src := `package service

// @auth "ActivateWorkflow" "workflow" {UserID: currentUser.ID, ResourceID: workflow.ID} "Forbidden" 401
func ActivateWorkflow() {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqAuth)
	assertEqual(t, "Action", seq.Action, "ActivateWorkflow")
	assertEqual(t, "Resource", seq.Resource, "workflow")
	assertEqual(t, "Message", seq.Message, "Forbidden")
	if seq.ErrStatus != 401 {
		t.Errorf("expected ErrStatus 401, got %d", seq.ErrStatus)
	}
}
