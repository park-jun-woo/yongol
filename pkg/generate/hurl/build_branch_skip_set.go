//ff:func feature=gen-hurl type=util control=sequence
//ff:what buildBranchSkipSet — BUG-016 Phase004: 분기 skip 해제, 모든 state transition을 smoke에 포함

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// buildBranchSkipSet returned the set of transition events that the
// smoke emitter should drop. Before Phase004 the implementation kept
// only the first outgoing transition per from-state so the smoke
// followed a single linear path — with the side effect that critical
// state-dependent endpoints (ExecuteWorkflow, ArchiveWorkflow) were
// removed from the scenario entirely (BUG-016).
//
// Phase004 makes smoke cover every operationId. Ordering concerns that
// used to justify skipping now live in buildTransitionOrder, which
// prioritises self-loops → non-terminal → terminal transitions so the
// linear walk stays executable even when several transitions share a
// from-state.
//
// The function is retained (as a no-op) so call sites in
// buildStateTransitions keep compiling; the return type stays identical
// and always yields an empty set. Inputs are ignored deliberately —
// deleting the parameters would ripple through callers without adding
// value.
func buildBranchSkipSet(diagrams []*statemachine.StateDiagram, transitionOrder map[string]int) map[string]bool {
	_ = diagrams
	_ = transitionOrder
	return map[string]bool{}
}
