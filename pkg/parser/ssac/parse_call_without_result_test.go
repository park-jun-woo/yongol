//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @call 결과 없음 파싱 검증 — Result nil, Inputs 확인

package ssac

import "testing"

func TestParseCallWithoutResult(t *testing.T) {
	src := `package service

// @call notification.Send({ID: reservation.ID, Status: "cancelled"})
func CancelReservation(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqCall)
	assertEqual(t, "Model", seq.Model, "notification.Send")
	if seq.Result != nil {
		t.Fatal("expected no result")
	}
	if len(seq.Inputs) != 2 {
		t.Fatalf("expected 2 inputs, got %d", len(seq.Inputs))
	}
	assertEqual(t, "Inputs.ID", seq.Inputs["ID"], "reservation.ID")
	assertEqual(t, "Inputs.Status", seq.Inputs["Status"], `"cancelled"`)
}
