//ff:func feature=gen-hurl type=util control=iteration dimension=3
//ff:what buildTransitionOrder — stateDiagram BFS로 event → order idx 계산
package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// buildTransitionOrder walks each stateDiagram BFS from InitialState
// and assigns an increasing index to each event in first-visit order.
func buildTransitionOrder(diagrams []*statemachine.StateDiagram) map[string]int {
	order := map[string]int{}
	idx := 0
	for _, d := range diagrams {
		if d == nil || d.InitialState == "" {
			continue
		}
		visited := map[string]bool{d.InitialState: true}
		queue := []string{d.InitialState}
		for len(queue) > 0 {
			state := queue[0]
			queue = queue[1:]
			for _, t := range d.Transitions {
				if t.From != state || visited[t.To] {
					continue
				}
				if _, exists := order[t.Event]; !exists {
					order[t.Event] = idx
					idx++
				}
				visited[t.To] = true
				queue = append(queue, t.To)
			}
		}
	}
	return order
}
