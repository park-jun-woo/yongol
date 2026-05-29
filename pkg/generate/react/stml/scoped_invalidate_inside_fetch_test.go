//ff:func feature=stml-gen type=test control=sequence
//ff:what fetch 블록 내 action의 scoped invalidation 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestScopedInvalidate_ActionInsideFetch(t *testing.T) {
	page, err := stmlparser.ParseReader("workflow-detail.html", strings.NewReader(`<main>
  <article data-fetch="GetWorkflow" data-param-id="route.id">
    <h2 data-bind="workflow.title"></h2>
    <footer data-state="workflow.status=draft">
      <button data-action="ActivateWorkflow" data-param-id="route.id">Activate</button>
    </footer>
  </article>
  <section data-fetch="ListActions" data-param-id="route.id">
    <ul data-each="actions">
      <li><span data-bind="action_type"></span></li>
    </ul>
  </section>
</main>`))
	if err != nil {
		t.Fatal(err)
	}
	code := GeneratePage(page, "")

	assertContains(t, code, "activateWorkflowMutation")
	assertContains(t, code, "queryKey: ['GetWorkflow']")

	mutIdx := strings.Index(code, "activateWorkflowMutation")
	if mutIdx < 0 {
		t.Fatal("activateWorkflowMutation not found")
	}
	onSuccessIdx := strings.Index(code[mutIdx:], "onSuccess")
	if onSuccessIdx < 0 {
		t.Fatal("onSuccess not found in activateWorkflowMutation")
	}
	closingIdx := strings.Index(code[mutIdx+onSuccessIdx:], "})")
	if closingIdx < 0 {
		t.Fatal("closing }) not found")
	}
	mutBlock := code[mutIdx : mutIdx+onSuccessIdx+closingIdx]
	if strings.Contains(mutBlock, "ListActions") {
		t.Errorf("activateWorkflowMutation should not invalidate ListActions, got:\n%s", mutBlock)
	}
}
