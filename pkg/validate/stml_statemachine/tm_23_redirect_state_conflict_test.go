//ff:func feature=validate type=test control=sequence topic=stml-statemachine
//ff:what TestTM23RedirectStateConflict — 도착 상태 충돌 WARNING·양립·침묵 케이스 검증

package stml_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM23RedirectStateConflict(t *testing.T) {
	diagrams := buildDiagramBySymbol([]*statemachine.StateDiagram{workflowDiagram()})

	boardPage := func(cond string) stml.PageSpec {
		return stml.PageSpec{
			FileName: "board.html",
			Children: []stml.ChildNode{{Kind: "state", State: &stml.StateBind{Condition: cond}}},
		}
	}
	action := stml.ActionBlock{OperationID: "ActivateWorkflow", Redirect: "/board"}

	// ActivateWorkflow arrives at "active" but the target page requires
	// "draft" → 1 WARNING.
	pages := []stml.PageSpec{boardPage("workflow.status=draft")}
	got := tm23RedirectStateConflict(action, "login.html", diagrams, pages)
	if len(got) != 1 || got[0].Level != diagnostic.LevelWarning {
		t.Fatalf("conflict: expected 1 WARNING, got %+v", got)
	}

	// Guard requires the arrival state → no diagnostics.
	pages = []stml.PageSpec{boardPage("workflow.status=active")}
	if d := tm23RedirectStateConflict(action, "login.html", diagrams, pages); len(d) != 0 {
		t.Errorf("compatible: expected 0 diagnostics, got %+v", d)
	}

	// Non-equality comparison is not comparable → silent.
	pages = []stml.PageSpec{boardPage("workflow.status!=active")}
	if d := tm23RedirectStateConflict(action, "login.html", diagrams, pages); len(d) != 0 {
		t.Errorf("non-equality: expected 0 diagnostics, got %+v", d)
	}

	// Syntactically invalid guard → ParseGuard errors, condition skipped.
	pages = []stml.PageSpec{boardPage("workflow.status=a &&")}
	if d := tm23RedirectStateConflict(action, "login.html", diagrams, pages); len(d) != 0 {
		t.Errorf("invalid guard: expected 0 diagnostics, got %+v", d)
	}

	// Guard on a model with no matching diagram → silent.
	pages = []stml.PageSpec{boardPage("invoice.status=draft")}
	if d := tm23RedirectStateConflict(action, "login.html", diagrams, pages); len(d) != 0 {
		t.Errorf("other model: expected 0 diagnostics, got %+v", d)
	}

	// Action that is no transition label → silent.
	pages = []stml.PageSpec{boardPage("workflow.status=draft")}
	other := stml.ActionBlock{OperationID: "CreateWorkflow", Redirect: "/board"}
	if d := tm23RedirectStateConflict(other, "login.html", diagrams, pages); len(d) != 0 {
		t.Errorf("non-transition: expected 0 diagnostics, got %+v", d)
	}

	// Redirect that resolves to no page → silent (TM-26 reports it).
	if d := tm23RedirectStateConflict(stml.ActionBlock{OperationID: "ActivateWorkflow", Redirect: "/nope"}, "login.html", diagrams, pages); len(d) != 0 {
		t.Errorf("unresolved redirect: expected 0 diagnostics, got %+v", d)
	}

	// No redirect → silent.
	if d := tm23RedirectStateConflict(stml.ActionBlock{OperationID: "ActivateWorkflow"}, "login.html", diagrams, pages); len(d) != 0 {
		t.Errorf("no redirect: expected 0 diagnostics, got %+v", d)
	}
}
