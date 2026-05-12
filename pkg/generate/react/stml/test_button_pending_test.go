//ff:func feature=stml-gen type=test control=sequence
//ff:what 액션 버튼과 제출 버튼의 isPending 로딩 상태를 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestActionButtonPending(t *testing.T) {
	page, _ := stmlparser.ParseReader("action-page.html", strings.NewReader(`<main>
  <button data-action="ActivateWorkflow">활성화</button>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, `disabled={activateWorkflowMutation.isPending}`)
	assertContains(t, code, `{activateWorkflowMutation.isPending ? '처리 중...' : '활성화'}`)
}

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

func TestActionButtonPendingWithWrapper(t *testing.T) {
	page, _ := stmlparser.ParseReader("action-page.html", strings.NewReader(`<main>
  <article data-fetch="GetReservation">
    <footer data-state="canCancel" class="mt-8">
      <button data-action="CancelReservation">예약 취소</button>
    </footer>
  </article>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, `disabled={cancelReservationMutation.isPending}`)
	assertContains(t, code, `{cancelReservationMutation.isPending ? '처리 중...' : '예약 취소'}`)
}
