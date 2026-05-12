//ff:func feature=stml-gen type=test control=sequence
//ff:what fetch/state 래퍼 내 액션 버튼의 isPending 로딩 상태를 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

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
