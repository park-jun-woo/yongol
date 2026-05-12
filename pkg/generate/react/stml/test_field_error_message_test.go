//ff:func feature=stml-gen type=test control=sequence
//ff:what 폼 필드 아래 formState.errors 에러 메시지 생성 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFieldErrorMessage(t *testing.T) {
	page, _ := stmlparser.ParseReader("form-page.html", strings.NewReader(`<main>
  <div data-action="CreateWorkflow">
    <input data-field="title" placeholder="제목" />
    <button type="submit">생성</button>
  </div>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, `{createWorkflowForm.formState.errors.title && (`)
	assertContains(t, code, `<p className="text-sm text-destructive">{createWorkflowForm.formState.errors.title?.message}</p>`)
}
