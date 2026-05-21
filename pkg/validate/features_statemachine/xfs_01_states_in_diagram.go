//ff:func feature=validate type=rule control=iteration dimension=2 topic=features-statemachine
//ff:what XFS-01 — features table의 state가 stateDiagram에 없으면 ERROR
package features_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xfs01StatesInDiagram validates XFS-01: for every table in FeatureTables
// that declares states, each state value must exist as a state in the
// corresponding stateDiagram.
func xfs01StatesInDiagram(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.FeatureTables == nil || fs.StateDiagrams == nil {
		return nil
	}
	diagramStates := buildDiagramStateMap(fs.StateDiagrams)
	var diags []diagnostic.Diagnostic
	for tableName, def := range fs.FeatureTables {
		if len(def.States) == 0 {
			continue
		}
		stateSet, hasDiagram := diagramStates[tableName]
		for _, state := range def.States {
			if hasDiagram && stateSet[state] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    "features.yaml",
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: `[XFS-01] table "` + tableName + `" declares state "` + state + `" but it is not in stateDiagram`,
			})
		}
	}
	return diags
}

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
