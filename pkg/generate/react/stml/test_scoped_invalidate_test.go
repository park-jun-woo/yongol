//ff:func feature=stml-gen type=test control=sequence
//ff:what mutation onSuccess가 소속 fetch 블록 기준으로 scoped invalidate하는지 검증
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestScopedInvalidate_ActionInsideFetch(t *testing.T) {
	// ActivateWorkflow is inside GetWorkflow fetch (via data-state) →
	// should only invalidate GetWorkflow, not ListActions etc.
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

	// ActivateWorkflow should invalidate GetWorkflow only
	assertContains(t, code, "activateWorkflowMutation")
	assertContains(t, code, "queryKey: ['GetWorkflow']")

	// ActivateWorkflow should NOT invalidate ListActions
	// Find the activateWorkflow mutation block and check it doesn't have ListActions
	mutIdx := strings.Index(code, "activateWorkflowMutation")
	if mutIdx < 0 {
		t.Fatal("activateWorkflowMutation not found")
	}
	// Find the onSuccess block for this mutation
	onSuccessIdx := strings.Index(code[mutIdx:], "onSuccess")
	if onSuccessIdx < 0 {
		t.Fatal("onSuccess not found in activateWorkflowMutation")
	}
	// Find the closing of this mutation (next "})") after onSuccess
	closingIdx := strings.Index(code[mutIdx+onSuccessIdx:], "})")
	if closingIdx < 0 {
		t.Fatal("closing }) not found")
	}
	mutBlock := code[mutIdx : mutIdx+onSuccessIdx+closingIdx]
	if strings.Contains(mutBlock, "ListActions") {
		t.Errorf("activateWorkflowMutation should not invalidate ListActions, got:\n%s", mutBlock)
	}
}

func TestScopedInvalidate_TopLevelAction(t *testing.T) {
	// CreateAction is at top level → should invalidate all fetches
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

	// CreateAction should invalidate both fetches
	assertContains(t, code, "createActionMutation")
	// Extract the full mutation block: from "createActionMutation" to the next
	// blank line or "return (" which marks the end of mutation declarations.
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

func TestScopedInvalidate_ActionInsideNestedState(t *testing.T) {
	// CancelReservation is inside GetReservation fetch (via data-state) →
	// should only invalidate GetReservation
	page, err := stmlparser.ParseReader("reservation-detail-page.html", strings.NewReader(`<main>
  <article data-fetch="GetReservation" data-param-reservation-id="route.ReservationID">
    <span data-bind="reservation.Status"></span>
    <footer data-state="canCancel">
      <button data-action="CancelReservation" data-param-reservation-id="route.ReservationID">Cancel</button>
    </footer>
  </article>
</main>`))
	if err != nil {
		t.Fatal(err)
	}
	code := GeneratePage(page, "")
	assertContains(t, code, "cancelReservationMutation")
	assertContains(t, code, "queryKey: ['GetReservation']")
}
