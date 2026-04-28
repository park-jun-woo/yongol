//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-statemachine
//ff:what transitionFrom — diagram 에서 op 이벤트의 From 상태 반환

package hurl_statemachine

import "github.com/park-jun-woo/yongol/pkg/parser/statemachine"

// transitionFrom returns the From state of the first transition whose
// label matches op. Empty string means the op is not part of this
// diagram.
func transitionFrom(d *statemachine.StateDiagram, op string) string {
	if d == nil {
		return ""
	}
	for _, t := range d.Transitions {
		if t.Event == op {
			return t.From
		}
	}
	return ""
}
