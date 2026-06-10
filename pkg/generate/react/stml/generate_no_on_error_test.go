//ff:func feature=stml-gen type=test control=sequence
//ff:what data-on-error 미선언 시 onError 핸들러·useState 미방출 검증 (기존 동작 불변)
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGeneratePage_NoOnError_NoHandler(t *testing.T) {
	page, _ := stmlparser.ParseReader("login-page.html", strings.NewReader(`<main>
  <div data-action="Login">
    <input data-field="Email" type="email" />
    <button type="submit">로그인</button>
  </div>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{APIImportPath: "@/lib/api"})

	assertNotContains(t, code, "onError:")
	assertNotContains(t, code, "useState")
}
