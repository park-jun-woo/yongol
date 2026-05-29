//ff:func feature=stml-gen type=test control=sequence
//ff:what GenerateOptions 기본값 적용 검증
package stml

import ("strings"; "testing"; stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml")

func TestGenerateOptionsDefaults(t *testing.T) {
	page, _ := stmlparser.ParseReader("login-page.html", strings.NewReader(`<main>
  <div data-action="Login">
    <input data-field="Email" type="email" />
    <button type="submit">로그인</button>
  </div>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, `import { api } from '@/lib/api'`)
	assertNotContains(t, code, "'use client'")
}
