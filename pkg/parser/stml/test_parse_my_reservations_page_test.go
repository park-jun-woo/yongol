//ff:func feature=stml-parse type=parser control=sequence
//ff:what TestParseMyReservationsPage — page with fetch/each/state blocks and one action

package stml

import (
	"strings"
	"testing"
)

func TestParseMyReservationsPage(t *testing.T) {
	input := `<main>
  <section data-fetch="ListMyReservations">
    <ul data-each="reservations">
      <li>
        <span data-bind="RoomID"></span>
        <span data-bind="StartAt"></span>
        <span data-bind="EndAt"></span>
        <span data-bind="Status"></span>
      </li>
    </ul>
    <p data-state="reservations.empty">예약이 없습니다</p>
  </section>

  <div data-action="CreateReservation">
    <input data-field="RoomID" type="number" />
    <div data-component="DatePicker" data-field="StartAt" />
    <div data-component="DatePicker" data-field="EndAt" />
    <button type="submit">예약하기</button>
  </div>
</main>`

	page, diags := ParseReader("my-reservations-page.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}

	// Fetch block
	if len(page.Fetches) != 1 {
		t.Fatalf("Fetches = %d, want 1", len(page.Fetches))
	}
	fetch := page.Fetches[0]
	if fetch.OperationID != "ListMyReservations" {
		t.Errorf("OperationID = %q, want %q", fetch.OperationID, "ListMyReservations")
	}

	// Each block
	if len(fetch.Eaches) != 1 {
		t.Fatalf("Eaches = %d, want 1", len(fetch.Eaches))
	}
	each := fetch.Eaches[0]
	if each.Field != "reservations" {
		t.Errorf("Each.Field = %q, want %q", each.Field, "reservations")
	}
	if len(each.Binds) != 4 {
		t.Errorf("Each.Binds = %d, want 4", len(each.Binds))
	}

	// State
	if len(fetch.States) != 1 {
		t.Fatalf("States = %d, want 1", len(fetch.States))
	}
	if fetch.States[0].Condition != "reservations.empty" {
		t.Errorf("State = %q, want %q", fetch.States[0].Condition, "reservations.empty")
	}

	// Action block
	if len(page.Actions) != 1 {
		t.Fatalf("Actions = %d, want 1", len(page.Actions))
	}
	action := page.Actions[0]
	if action.OperationID != "CreateReservation" {
		t.Errorf("OperationID = %q, want %q", action.OperationID, "CreateReservation")
	}
	if len(action.Fields) != 3 {
		t.Fatalf("Fields = %d, want 3", len(action.Fields))
	}
	assertField(t, action.Fields[0], "RoomID", "input", "number")
	assertField(t, action.Fields[1], "StartAt", "data-component:DatePicker", "")
	assertField(t, action.Fields[2], "EndAt", "data-component:DatePicker", "")
}
