//ff:func feature=stml-parse type=parser control=sequence
//ff:what TestPhase4_StateWithAction — state element containing a nested action child

package stml

import (
	"strings"
	"testing"
)

func TestPhase4_StateWithAction(t *testing.T) {
	input := `<main>
  <article data-fetch="GetReservation" data-param-reservation-id="route.ReservationID">
    <footer data-state="canCancel" class="mt-8 pt-4 border-t">
      <button data-action="CancelReservation" data-param-reservation-id="route.ReservationID">
        예약 취소
      </button>
    </footer>
  </article>
</main>`

	page, diags := ParseReader("test.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}

	state := page.Fetches[0].States[0]
	if state.Tag != "footer" {
		t.Errorf("State.Tag = %q, want %q", state.Tag, "footer")
	}
	if state.ClassName != "mt-8 pt-4 border-t" {
		t.Errorf("State.ClassName = %q, want %q", state.ClassName, "mt-8 pt-4 border-t")
	}
	if len(state.Children) != 1 {
		t.Fatalf("State.Children = %d, want 1", len(state.Children))
	}
	if state.Children[0].Kind != "action" {
		t.Errorf("State.Children[0].Kind = %q, want %q", state.Children[0].Kind, "action")
	}
	if state.Children[0].Action.OperationID != "CancelReservation" {
		t.Errorf("Action.OperationID = %q, want %q", state.Children[0].Action.OperationID, "CancelReservation")
	}
}
