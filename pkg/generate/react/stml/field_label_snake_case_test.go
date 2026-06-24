//ff:func feature=stml-gen type=test control=sequence
//ff:what 폼 필드 snake_case label + id 속성 자동 생성 검증
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
	assertContains(t, code, `<label htmlFor="createWorkflow-trigger_event" className="text-sm font-medium">Trigger Event</label>`)
	assertContains(t, code, `id="createWorkflow-trigger_event"`)
}
