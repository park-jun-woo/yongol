//ff:func feature=validate type=helper control=iteration dimension=1 topic=features-statemachine
//ff:what buildDiagramStateMap — stateDiagram 배열을 ID → state 이름 집합 맵으로 변환

package features_statemachine

import "github.com/park-jun-woo/yongol/pkg/parser/statemachine"

// buildDiagramStateMap creates a lookup: diagram ID → set of state names.
func buildDiagramStateMap(diagrams []*statemachine.StateDiagram) map[string]map[string]bool {
	m := make(map[string]map[string]bool, len(diagrams))
	for _, d := range diagrams {
		set := make(map[string]bool, len(d.States))
		for _, s := range d.States {
			set[s] = true
		}
		m[d.ID] = set
	}
	return m
}
