//ff:func feature=stml-gen type=test control=sequence
//ff:what 제출 버튼의 isPending 로딩 상태를 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSubmitButtonPending(t *testing.T) {
	page, _ := stmlparser.ParseReader("form-page.html", strings.NewReader(`<main>
  <div data-action="Login">
    <input data-field="Email" type="email" placeholder="이메일" />
    <button type="submit">로그인</button>
  </div>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, `disabled={loginMutation.isPending}`)
	assertContains(t, code, `{loginMutation.isPending ? '처리 중...' : '로그인'}`)
	assertContains(t, code, `type="submit"`)
}
