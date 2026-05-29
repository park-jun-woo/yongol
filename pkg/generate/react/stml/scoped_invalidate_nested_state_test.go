//ff:func feature=stml-gen type=test control=sequence
//ff:what nested state 내 action의 scoped invalidation 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestScopedInvalidate_ActionInsideNestedState(t *testing.T) {
	page, err := stmlparser.ParseReader("reservation-detail-page.html", strings.NewReader(`<main>
  <article data-fetch="GetReservation" data-param-reservation-id="route.ReservationID">
    <span data-bind="reservation.Status"></span>
    <footer data-state="canCancel">
      <button data-action="CancelReservation" data-param-reservation-id="route.ReservationID">Cancel</button>
    </footer>
  </article>
</main>`))
	if err != nil {
		t.Fatal(err)
	}
	code := GeneratePage(page, "")
	assertContains(t, code, "cancelReservationMutation")
	assertContains(t, code, "queryKey: ['GetReservation']")
}
