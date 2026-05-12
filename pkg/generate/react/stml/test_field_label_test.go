//ff:func feature=stml-gen type=test control=sequence
//ff:what 폼 필드에 label + id 속성 자동 생성을 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFieldLabelSnakeCase(t *testing.T) {
	page, _ := stmlparser.ParseReader("form-page.html", strings.NewReader(`<main>
  <div data-action="CreateWorkflow">
    <input data-field="trigger_event" placeholder="이벤트" />
    <button type="submit">생성</button>
  </div>
</main>`))
	code := GeneratePage(page, "")
	// snake_case → Title Case label
	assertContains(t, code, `<label htmlFor="trigger_event">Trigger Event</label>`)
	assertContains(t, code, `id="trigger_event"`)
}

func TestFieldLabelPascalCase(t *testing.T) {
	page, _ := stmlparser.ParseReader("form-page.html", strings.NewReader(`<main>
  <div data-action="UpdateRoom">
    <input data-field="RoomID" placeholder="ID" />
    <button type="submit">수정</button>
  </div>
</main>`))
	code := GeneratePage(page, "")
	// PascalCase → spaced label
	assertContains(t, code, `<label htmlFor="RoomID">Room ID</label>`)
	assertContains(t, code, `id="RoomID"`)
}

func TestFieldLabelComponentNoLabel(t *testing.T) {
	page, _ := stmlparser.ParseReader("form-page.html", strings.NewReader(`<main>
  <div data-action="CreateReservation">
    <div data-component="DatePicker" data-field="StartAt" />
    <button type="submit">예약</button>
  </div>
</main>`))
	code := GeneratePage(page, "")
	// data-component fields should NOT get label wrapper
	assertNotContains(t, code, `<label htmlFor="StartAt">`)
	assertContains(t, code, "<DatePicker")
}
