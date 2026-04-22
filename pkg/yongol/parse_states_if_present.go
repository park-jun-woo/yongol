//ff:func feature=orchestrator type=loader control=sequence
//ff:what States(Mermaid stateDiagram) 탐지 시 ParseDir 실행
package yongol

import (
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// parseStatesIfPresent parses the state-machine directory when present.
func parseStatesIfPresent(fs *Fullstack, has map[SSOTKind]DetectedSSOT) {
	d, ok := has[KindStates]
	if !ok {
		return
	}
	diagrams, diags := statemachine.ParseDir(d.Path)
	fs.ParseDiagnostics = append(fs.ParseDiagnostics, diags...)
	if len(diags) == 0 {
		fs.StateDiagrams = diagrams
	}
}
