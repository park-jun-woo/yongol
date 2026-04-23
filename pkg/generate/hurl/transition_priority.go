//ff:func feature=gen-hurl type=util control=selection
//ff:what transitionPriority — 단일 transition의 우선순위 점수 (0=self-loop / 1=비종착 / 2=종착)

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// transitionPriority returns the sort key used by outgoingByPriority.
// Lower score emits earlier: self-loops (0) before non-terminal (1)
// before terminal (2). The tiers keep smoke scenarios executable —
// a terminal transition typically locks the record so everything that
// depends on an editable state must run first.
func transitionPriority(t statemachine.Transition, fromState string, terminals map[string]bool) int {
	switch {
	case t.To == fromState:
		return 0
	case terminals[t.To]:
		return 2
	default:
		return 1
	}
}
