//ff:func feature=stml-parse type=parser control=sequence
//ff:what TestParseRoomEditPage — edit page with two actions (update + delete) and route params

package stml

import (
	"strings"
	"testing"
)

func TestParseRoomEditPage(t *testing.T) {
	input := `<main>
  <div data-action="UpdateRoom" data-param-room-id="route.RoomID">
    <input data-field="Name" />
    <input data-field="Capacity" type="number" />
    <input data-field="Location" />
    <button type="submit">수정</button>
  </div>

  <footer data-state="canDelete">
    <button data-action="DeleteRoom" data-param-room-id="route.RoomID">
      스터디룸 삭제
    </button>
  </footer>
</main>`

	page, diags := ParseReader("room-edit-page.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}

	if len(page.Actions) != 2 {
		t.Fatalf("Actions = %d, want 2", len(page.Actions))
	}

	update := page.Actions[0]
	if update.OperationID != "UpdateRoom" {
		t.Errorf("OperationID = %q, want %q", update.OperationID, "UpdateRoom")
	}
	if len(update.Params) != 1 {
		t.Fatalf("Params = %d, want 1", len(update.Params))
	}
	if update.Params[0].Name != "roomId" {
		t.Errorf("Param.Name = %q, want %q", update.Params[0].Name, "roomId")
	}
	if len(update.Fields) != 3 {
		t.Fatalf("Fields = %d, want 3", len(update.Fields))
	}

	del := page.Actions[1]
	if del.OperationID != "DeleteRoom" {
		t.Errorf("OperationID = %q, want %q", del.OperationID, "DeleteRoom")
	}
	if len(del.Params) != 1 {
		t.Fatalf("Params = %d, want 1", len(del.Params))
	}
}
