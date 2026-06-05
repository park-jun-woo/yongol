//ff:func feature=validate type=helper control=iteration dimension=1 topic=stml-statemachine
//ff:what buildDiagramBySymbol — stateDiagram 배열을 Symbol(케이스 안정 키) → *StateDiagram 맵으로 변환

package stml_statemachine

import "github.com/park-jun-woo/yongol/pkg/parser/statemachine"

// buildDiagramBySymbol creates a lookup: diagram Symbol → diagram. Symbol is the
// PascalCase form of the filename stem, so STML guard model prefixes can match
// after normalization (modelSymbol).
func buildDiagramBySymbol(diagrams []*statemachine.StateDiagram) map[string]*statemachine.StateDiagram {
	m := make(map[string]*statemachine.StateDiagram, len(diagrams))
	for _, d := range diagrams {
		m[d.Symbol] = d
	}
	return m
}
