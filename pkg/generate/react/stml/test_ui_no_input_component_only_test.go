//ff:func feature=stml-gen type=test control=sequence
//ff:what data-component 전용 폼에서 Input import 미생성을 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestUINoInputImportComponentOnly(t *testing.T) {
	page, _ := stmlparser.ParseReader("form-page.html", strings.NewReader(`<main>
  <div data-action="CreateReservation">
    <div data-component="DatePicker" data-field="StartAt" />
    <button type="submit">예약</button>
  </div>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, "import { Button } from '@/components/ui/Button'")
	assertNotContains(t, code, "import { Input }")
}
