//ff:func feature=stml-gen type=test control=sequence
//ff:what 폼 필드 PascalCase label + id 속성 자동 생성 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFieldLabelPascalCase(t *testing.T) {
	page, _ := stmlparser.ParseReader("form-page.html", strings.NewReader(`<main>
  <div data-action="UpdateRoom">
    <input data-field="RoomID" placeholder="ID" />
    <button type="submit">수정</button>
  </div>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, `<label htmlFor="RoomID">Room ID</label>`)
	assertContains(t, code, `id="RoomID"`)
}
