//ff:func feature=ssac-parse type=parser control=sequence
//ff:what 전체 시퀀스 통합 파싱 검증 — auth/get/empty/state/call/put/get/response 8단계

package ssac

import "testing"

func TestParseFullExample(t *testing.T) {
	src := `package service

import "myapp/auth"

// @auth "cancel" "reservation" {id: request.ReservationID} "권한 없음"
// @get Reservation reservation = Reservation.FindByID({ReservationID: request.ReservationID})
// @empty reservation "예약을 찾을 수 없습니다"
// @state reservation {status: reservation.Status} "cancel" "취소할 수 없습니다"
// @call Refund refund = billing.CalculateRefund({ID: reservation.ID, StartAt: reservation.StartAt, EndAt: reservation.EndAt})
// @put Reservation.UpdateStatus({ReservationID: request.ReservationID, Status: "cancelled"})
// @get Reservation reservation = Reservation.FindByID({ReservationID: request.ReservationID})
// @response {
//   reservation: reservation,
//   refund: refund
// }
func CancelReservation(c *gin.Context) {}
`
	sfs := parseTestFile(t, src)
	sf := sfs[0]
	assertEqual(t, "Name", sf.Name, "CancelReservation")

	if len(sf.Sequences) != 8 {
		t.Fatalf("expected 8 sequences, got %d", len(sf.Sequences))
	}

	// @auth
	assertEqual(t, "seq0.Type", sf.Sequences[0].Type, SeqAuth)
	assertEqual(t, "seq0.Action", sf.Sequences[0].Action, "cancel")

	// @get
	assertEqual(t, "seq1.Type", sf.Sequences[1].Type, SeqGet)
	assertEqual(t, "seq1.Model", sf.Sequences[1].Model, "Reservation.FindByID")

	// @empty
	assertEqual(t, "seq2.Type", sf.Sequences[2].Type, SeqEmpty)

	// @state
	assertEqual(t, "seq3.Type", sf.Sequences[3].Type, SeqState)
	assertEqual(t, "seq3.DiagramID", sf.Sequences[3].DiagramID, "reservation")

	// @call
	assertEqual(t, "seq4.Type", sf.Sequences[4].Type, SeqCall)
	assertEqual(t, "seq4.Model", sf.Sequences[4].Model, "billing.CalculateRefund")
	if seq4r := sf.Sequences[4].Result; seq4r == nil {
		t.Fatal("expected call result")
	} else {
		assertEqual(t, "seq4.Result.Type", seq4r.Type, "Refund")
	}

	// @put
	assertEqual(t, "seq5.Type", sf.Sequences[5].Type, SeqPut)

	// @get (re-fetch)
	assertEqual(t, "seq6.Type", sf.Sequences[6].Type, SeqGet)

	// @response
	assertEqual(t, "seq7.Type", sf.Sequences[7].Type, SeqResponse)
	assertEqual(t, "seq7.Fields[reservation]", sf.Sequences[7].Fields["reservation"], "reservation")
	assertEqual(t, "seq7.Fields[refund]", sf.Sequences[7].Fields["refund"], "refund")
}
