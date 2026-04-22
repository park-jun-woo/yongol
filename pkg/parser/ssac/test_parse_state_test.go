//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @state 파싱 검증 — DiagramID, Transition, Message, Inputs 확인

package ssac

import "testing"

func TestParseState(t *testing.T) {
	src := `package service

// @state reservation {status: reservation.Status} "cancel" "취소할 수 없습니다"
func CancelReservation(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqState)
	assertEqual(t, "DiagramID", seq.DiagramID, "reservation")
	assertEqual(t, "Transition", seq.Transition, "cancel")
	assertEqual(t, "Message", seq.Message, "취소할 수 없습니다")
	if seq.Inputs["status"] != "reservation.Status" {
		t.Errorf("expected Inputs[status]=%q, got %q", "reservation.Status", seq.Inputs["status"])
	}
}
