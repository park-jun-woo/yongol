//ff:func feature=stml-gen type=test control=sequence
//ff:what NoBodyOps에 미포함된 액션이 mutate({})를 생성하는지 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestActionButtonNonVoidMutate(t *testing.T) {
	page, _ := stmlparser.ParseReader("activate-page.html", strings.NewReader(`<main>
  <button data-action="ActivateWorkflow">활성화</button>
</main>`))
	code := GeneratePage(page, "")
	assertContains(t, code, `activateWorkflowMutation.mutate({})`)
}
