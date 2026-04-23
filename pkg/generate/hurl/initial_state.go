//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what initialState — 첫 번째 stateDiagram의 InitialState 반환 (비어있으면 "")

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// initialState returns the initial state of the first non-empty diagram.
// Used as the seed state for the state-transition walker.
func initialState(diagrams []*statemachine.StateDiagram) string {
	for _, d := range diagrams {
		if d == nil {
			continue
		}
		if d.InitialState != "" {
			return d.InitialState
		}
	}
	return ""
}
