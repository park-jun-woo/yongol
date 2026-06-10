//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-statemachine
//ff:what TM-23 — data-redirect 대상 페이지의 상태 가드가 액션의 stateDiagram 전이 도착 상태와 양립 불가 (WARNING)

package stml_statemachine

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm23RedirectStateConflict is the runtime twin of XOH-05 (call order must
// satisfy the stateDiagram). When the action's OperationID is a transition
// label, the diagram yields the set S of arrival states; if the redirect
// target page has a page-level data-state guard comparing the same
// diagram's state with "=" and the required state is not in S, the user
// lands on a screen whose guard can never hold right after the action.
// Anything not comparable — unresolved redirect (TM-26 reports it), a
// different model, a non-equality or negated/OR guard, no guard at all,
// or an OperationID that is no transition label — stays silent (WARNING
// by design: intentional exceptions are allowed).
func tm23RedirectStateConflict(a stml.ActionBlock, file string, diagramBySymbol map[string]*statemachine.StateDiagram, pages []stml.PageSpec) []diagnostic.Diagnostic {
	if a.Redirect == "" {
		return nil
	}
	target := findRedirectTargetPage(a.Redirect, pages)
	if target == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, cond := range collectPageStateConditions(target.Children) {
		expr, err := stml.ParseGuard(cond)
		if err != nil {
			continue
		}
		for _, pair := range collectEqualComparePairs(expr) {
			d, ok := diagramBySymbol[modelSymbol(pair.Model)]
			if !ok {
				continue
			}
			arrival := transitionToStates(d, a.OperationID)
			if len(arrival) == 0 {
				continue
			}
			if stateInSlice(arrival, pair.Value) {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:        file,
				Phase:       diagnostic.PhaseValidate,
				Level:       diagnostic.LevelWarning,
				Message:     fmt.Sprintf("[TM-23] action %q redirects to %q, but that page's data-state guard requires %q to be %q, which is not an arrival state of transition %q in stateDiagram %q", a.OperationID, a.Redirect, pair.Model, pair.Value, a.OperationID, d.Symbol),
				Advice:      fmt.Sprintf("Redirect to a page whose guard matches an arrival state of %q (%s), or adjust the guard or the stateDiagram", a.OperationID, strings.Join(arrival, ", ")),
				OperationID: a.OperationID,
			})
		}
	}
	return diags
}
