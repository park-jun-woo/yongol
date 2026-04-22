//ff:func feature=ssac-parse type=parser control=sequence
//ff:what @delete 파싱 검증 — 결과 없음, Inputs 확인

package ssac

import "testing"

func TestParseDelete(t *testing.T) {
	src := `package service

// @delete Reservation.Cancel({ID: reservation.ID})
func CancelReservation(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Type", seq.Type, SeqDelete)
	if seq.Result != nil {
		t.Fatal("expected no result for @delete")
	}
	assertEqual(t, "Inputs.ID", seq.Inputs["ID"], "reservation.ID")
}
