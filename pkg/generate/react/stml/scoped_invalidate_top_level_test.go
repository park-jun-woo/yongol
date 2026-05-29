//ff:func feature=stml-gen type=test control=sequence
//ff:what top-level action의 전체 fetch invalidation 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestScopedInvalidate_TopLevelAction(t *testing.T) {
	page, err := stmlparser.ParseReader("workflow-detail.html", strings.NewReader(`<main>
  <article data-fetch="GetWorkflow" data-param-id="route.id">
    <h2 data-bind="workflow.title"></h2>
  </article>
  <section data-fetch="ListActions" data-param-id="route.id">
    <ul data-each="actions">
      <li><span data-bind="action_type"></span></li>
    </ul>
  </section>
  <div data-action="CreateAction" data-param-id="route.id">
    <input data-field="action_type" type="text" placeholder="Type" />
    <button type="submit">Add</button>
  </div>
</main>`))
	if err != nil {
		t.Fatal(err)
	}
	code := GeneratePage(page, "")

	assertContains(t, code, "createActionMutation")
	mutIdx := strings.Index(code, "createActionMutation")
	if mutIdx < 0 {
		t.Fatal("createActionMutation not found")
	}
	rest := code[mutIdx:]
	endIdx := strings.Index(rest, "\n\n")
	if endIdx < 0 {
		endIdx = len(rest)
	}
	mutBlock := rest[:endIdx]
	if !strings.Contains(mutBlock, "GetWorkflow") {
		t.Errorf("createActionMutation should invalidate GetWorkflow, got:\n%s", mutBlock)
	}
	if !strings.Contains(mutBlock, "ListActions") {
		t.Errorf("createActionMutation should invalidate ListActions, got:\n%s", mutBlock)
	}
}
