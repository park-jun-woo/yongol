//ff:func feature=ssac-parse type=parser control=sequence
//ff:what 리터럴 문자열 인자 파싱 검증 — 따옴표 포함 값 확인

package ssac

import "testing"

func TestParseLiteralArg(t *testing.T) {
	src := `package service

// @put Reservation.UpdateStatus({ReservationID: request.ReservationID, Status: "cancelled"})
func CancelReservation(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	seq := sfs[0].Sequences[0]
	assertEqual(t, "Inputs.Status", seq.Inputs["Status"], `"cancelled"`)
}
