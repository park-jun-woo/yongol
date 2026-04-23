//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what outgoingByPriority — 한 상태의 outgoing transitions을 self-loop→비종착→종착 순으로 정렬

package hurl

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// outgoingByPriority returns the transitions outgoing from `state`
// sorted by execution priority:
//
//  1. Self-loops (t.To == state) — state preserving, emit first.
//  2. Transitions to non-terminal states — follow-on steps remain valid.
//  3. Transitions to terminal states — one-way, emit last.
//
// Within each bucket, transitions preserve their original diagram order
// (sort is stable) so authors can influence the sequence by how they
// arrange transitions in the Mermaid source.
func outgoingByPriority(d *statemachine.StateDiagram, state string, terminals map[string]bool) []statemachine.Transition {
	var outs []statemachine.Transition
	for _, t := range d.Transitions {
		if t.From == state {
			outs = append(outs, t)
		}
	}
	sort.SliceStable(outs, func(i, j int) bool {
		return transitionPriority(outs[i], state, terminals) < transitionPriority(outs[j], state, terminals)
	})
	return outs
}
