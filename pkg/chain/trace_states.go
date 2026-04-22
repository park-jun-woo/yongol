//ff:func feature=chain type=util control=iteration dimension=1
//ff:what traceStates finds state diagrams referenced by @state sequences.
package chain

import (
	"log/slog"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

func traceStates(sf *ssac.ServiceFunc, diagrams []*statemachine.StateDiagram, specsDir string) []Link {
	diagramIDs := map[string]bool{}
	transitions := map[string]string{} // diagramID -> transition name
	for _, seq := range sf.Sequences {
		if seq.Type != "state" {
			continue
		}
		diagramIDs[seq.DiagramID] = true
		transitions[seq.DiagramID] = seq.Transition
	}

	if len(diagramIDs) == 0 {
		slog.Debug("chain.traceStates: no @state sequences in SSaC function", "operationId", sf.Name)
		return nil
	}

	var links []Link
	for _, d := range diagrams {
		if !diagramIDs[d.ID] {
			continue
		}
		links = append(links, buildStateLink(d.ID, specsDir, transitions))
	}
	return links
}
