//ff:func feature=validate type=util control=iteration dimension=1 topic=states
//ff:what StateDiagram 리스트를 ID→Diagram 맵으로 변환

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// buildDiagramByID builds a lookup map of diagrams by ID.
func buildDiagramByID(diagrams []*statemachine.StateDiagram) map[string]*statemachine.StateDiagram {
	m := make(map[string]*statemachine.StateDiagram, len(diagrams))
	for _, d := range diagrams {
		m[d.ID] = d
	}
	return m
}
