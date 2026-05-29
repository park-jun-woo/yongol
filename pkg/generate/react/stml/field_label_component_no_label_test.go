//ff:func feature=stml-gen type=test control=sequence
//ff:what data-component 필드에 label 미생성 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFieldLabelComponentNoLabel(t *testing.T) {
	page, _ := stmlparser.ParseReader("form-page.html", strings.NewReader(`<main>
  <div data-action="CreateReservation">
    <div data-component="DatePicker" data-field="StartAt" />
    <button type="submit">예약</button>
  </div>
</main>`))
	code := GeneratePage(page, "")
	assertNotContains(t, code, `<label htmlFor="StartAt">`)
	assertContains(t, code, "<DatePicker")
}
