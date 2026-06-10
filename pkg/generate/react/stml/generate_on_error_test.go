//ff:func feature=stml-gen type=test control=sequence
//ff:what data-on-error 선언 시 에러 상태 + onError 핸들러 + 조건부 렌더 방출 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGeneratePage_OnError_StateHandlerConditionalRender(t *testing.T) {
	page, _ := stmlparser.ParseReader("login-page.html", strings.NewReader(`<main>
  <div data-action="Login" data-redirect="/">
    <input data-field="Email" type="email" />
    <button type="submit">로그인</button>
    <p class="error" data-on-error></p>
  </div>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		APIImportPath: "@/lib/api",
		BearerAuth:    true,
	})

	// error message state + useState import
	assertContains(t, code, "const [loginError, setLoginError] = useState<string | null>(null)")
	assertContains(t, code, "import { useState } from 'react'")

	// onError handler feeds the state; onMutate clears it on (re)submission.
	// The thrown ErrorResponse is a plain object — message is extracted
	// defensively (no schema guarantee; XOE-01 checks only error/code).
	assertContains(t, code, "onError: (err) => {")
	assertContains(t, code, "const msg = (err as any)?.message")
	assertContains(t, code, "setLoginError(typeof msg === 'string' && msg !== '' ? msg : String(err))")
	assertContains(t, code, "onMutate: () => setLoginError(null)")

	// the data-on-error element renders conditionally with the message bound
	assertContains(t, code, `{loginError && <p className="error">{loginError}</p>}`)

	// the declared element owns the display — no default slot alongside it
	// (page-flow Phase004 emits the default only when data-on-error is absent)
	assertNotContains(t, code, `role="alert"`)
}
