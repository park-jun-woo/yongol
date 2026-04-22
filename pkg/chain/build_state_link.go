//ff:func feature=chain type=util control=sequence
//ff:what 단일 stateDiagram에서 Link를 생성한다
package chain

import "path/filepath"

// buildStateLink creates a Link for a matched state diagram.
func buildStateLink(diagramID, specsDir string, transitions map[string]string) Link {
	relPath := "states/" + diagramID + ".md"
	trans := transitions[diagramID]
	line := 0
	if trans != "" {
		line = grepLine(filepath.Join(specsDir, relPath), trans)
	}
	summary := "diagram: " + diagramID
	if trans != "" {
		summary += " -> " + trans
	}
	return Link{
		Kind:    "StateDiag",
		File:    relPath,
		Line:    line,
		Summary: summary,
	}
}
