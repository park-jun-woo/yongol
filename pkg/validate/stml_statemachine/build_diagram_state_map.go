//ff:func feature=validate type=helper control=iteration dimension=1 topic=stml-statemachine
//ff:what buildDiagramStateMap — stateDiagram 배열을 Symbol(케이스 안정 키) → state 이름 집합 맵으로 변환

package stml_statemachine

import "github.com/park-jun-woo/yongol/pkg/parser/statemachine"

// buildDiagramStateMap creates a lookup: diagram Symbol → set of state names.
// It mirrors features_statemachine.buildDiagramStateMap (which is unexported and
// keys on ID); this copy keys on Symbol (PascalCase) so STML guard model
// prefixes can match case-insensitively after normalization.
func buildDiagramStateMap(diagrams []*statemachine.StateDiagram) map[string]map[string]bool {
	m := make(map[string]map[string]bool, len(diagrams))
	for _, d := range diagrams {
		set := make(map[string]bool, len(d.States))
		for _, s := range d.States {
			set[s] = true
		}
		m[d.Symbol] = set
	}
	return m
}
