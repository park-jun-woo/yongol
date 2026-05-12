//ff:func feature=stml-gen type=test control=sequence
//ff:what 액션 버튼의 isPending 로딩 상태를 검증
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
