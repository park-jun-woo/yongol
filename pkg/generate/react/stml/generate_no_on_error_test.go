//ff:func feature=stml-gen type=test control=sequence
//ff:what data-on-error 미선언 시 기본 onError·에러 상태·기본 표시 슬롯 방출 검증 (Phase004 무음 실패 차단)
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestGeneratePage_NoOnError_DefaultEmission(t *testing.T) {
	page, _ := stmlparser.ParseReader("login-page.html", strings.NewReader(`<main>
  <div data-action="Login">
    <input data-field="Email" type="email" />
    <button type="submit">로그인</button>
  </div>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{APIImportPath: "@/lib/api"})

	// page-flow Phase004: without data-on-error the error state, the onError
	// handler and the default display slot are still emitted — a rejected
	// mutation must never look like a success (BUG-113 (2)).
	assertContains(t, code, "const [loginError, setLoginError] = useState<string | null>(null)")
	assertContains(t, code, "import { useState } from 'react'")
	assertContains(t, code, "onError: (err) => {")

	// the default slot renders right after the submit button
	assertContains(t, code, `{loginError && <p role="alert" className="text-sm text-destructive">{loginError}</p>}`)

	// resubmission clears the previous message before the request fires
	assertContains(t, code, "onMutate: () => setLoginError(null)")
}
