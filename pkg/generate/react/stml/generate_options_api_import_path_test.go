//ff:func feature=stml-gen type=test control=sequence
//ff:what GenerateOptions APIImportPath 커스텀 경로 검증
package stml

import ("strings"; "testing"; stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml")

func TestGenerateOptionsAPIImportPath(t *testing.T) {
	page, _ := stmlparser.ParseReader("login-page.html", strings.NewReader(`<main>
  <div data-action="Login">
    <input data-field="Email" type="email" />
    <button type="submit">로그인</button>
  </div>
</main>`))
	opts := GenerateOptions{APIImportPath: "../api", UseClient: false}
	code := GeneratePage(page, "", opts)
	assertContains(t, code, `import { api } from '../api'`)
	assertNotContains(t, code, `@/lib/api`)
	assertNotContains(t, code, "'use client'")
}
