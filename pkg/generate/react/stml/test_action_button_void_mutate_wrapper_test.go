//ff:func feature=stml-gen type=test control=sequence
//ff:what wrapper 내부 NoBodyOps 액션이 mutate()를 생성하는지 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestActionButtonVoidMutateWithWrapper(t *testing.T) {
	page, _ := stmlparser.ParseReader("detail-page.html", strings.NewReader(`<main>
  <article data-fetch="GetReservation">
    <footer data-state="canCancel" class="mt-8">
      <button data-action="CancelReservation">예약 취소</button>
    </footer>
  </article>
</main>`))
	code := GeneratePage(page, "", GenerateOptions{
		NoBodyOps: map[string]bool{"CancelReservation": true},
	})
	assertContains(t, code, `cancelReservationMutation.mutate()`)
	assertNotContains(t, code, `mutate({})`)
}
