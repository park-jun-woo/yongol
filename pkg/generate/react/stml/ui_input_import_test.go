//ff:func feature=stml-gen type=test control=sequence
//ff:what 폼 페이지에서 Input UI 컴포넌트 import 생성을 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestUIInputImport(t *testing.T) {
	page, _ := stmlparser.ParseReader("login-page.html", strings.NewReader(`<main>
  <div data-action="Login">
    <input data-field="Email" type="email" placeholder="이메일" />
    <button type="submit">로그인</button>
  </div>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, "import { Input } from '@/components/ui/Input'")
	assertContains(t, code, "<Input ")
	assertNotContains(t, code, "<input ")
}
