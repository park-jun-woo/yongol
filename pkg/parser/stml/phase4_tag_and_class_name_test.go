//ff:func feature=stml-parse type=parser control=sequence
//ff:what TestPhase4_TagAndClassName — tag, className, placeholder, submitText extraction

package stml

import (
	"strings"
	"testing"
)

func TestPhase4_TagAndClassName(t *testing.T) {
	input := `<main class="max-w-4xl mx-auto p-6">
  <section data-fetch="ListMyReservations" class="mb-8">
    <ul data-each="reservations" class="space-y-3">
      <li class="flex justify-between p-4 border rounded">
        <span data-bind="RoomID" class="font-semibold"></span>
      </li>
    </ul>
  </section>
  <div data-action="CreateReservation" class="space-y-4">
    <input data-field="RoomID" type="number" placeholder="스터디룸 번호" class="w-full px-3 py-2 border rounded" />
    <button type="submit">예약하기</button>
  </div>
</main>`

	page, diags := ParseReader("test.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}

	// Fetch Tag + ClassName
	fetch := page.Fetches[0]
	if fetch.Tag != "section" {
		t.Errorf("Fetch.Tag = %q, want %q", fetch.Tag, "section")
	}
	if fetch.ClassName != "mb-8" {
		t.Errorf("Fetch.ClassName = %q, want %q", fetch.ClassName, "mb-8")
	}

	// Each Tag + ClassName + ItemTag + ItemClassName
	each := fetch.Eaches[0]
	if each.Tag != "ul" {
		t.Errorf("Each.Tag = %q, want %q", each.Tag, "ul")
	}
	if each.ClassName != "space-y-3" {
		t.Errorf("Each.ClassName = %q, want %q", each.ClassName, "space-y-3")
	}
	if each.ItemTag != "li" {
		t.Errorf("Each.ItemTag = %q, want %q", each.ItemTag, "li")
	}
	if each.ItemClassName != "flex justify-between p-4 border rounded" {
		t.Errorf("Each.ItemClassName = %q, want %q", each.ItemClassName, "flex justify-between p-4 border rounded")
	}

	// Bind ClassName
	if each.Binds[0].ClassName != "font-semibold" {
		t.Errorf("Bind.ClassName = %q, want %q", each.Binds[0].ClassName, "font-semibold")
	}

	// Action Tag + ClassName
	action := page.Actions[0]
	if action.Tag != "div" {
		t.Errorf("Action.Tag = %q, want %q", action.Tag, "div")
	}
	if action.ClassName != "space-y-4" {
		t.Errorf("Action.ClassName = %q, want %q", action.ClassName, "space-y-4")
	}

	// Field Placeholder + ClassName
	field := action.Fields[0]
	if field.Placeholder != "스터디룸 번호" {
		t.Errorf("Field.Placeholder = %q, want %q", field.Placeholder, "스터디룸 번호")
	}
	if field.ClassName != "w-full px-3 py-2 border rounded" {
		t.Errorf("Field.ClassName = %q, want %q", field.ClassName, "w-full px-3 py-2 border rounded")
	}

	// SubmitText
	if action.SubmitText != "예약하기" {
		t.Errorf("SubmitText = %q, want %q", action.SubmitText, "예약하기")
	}
}
