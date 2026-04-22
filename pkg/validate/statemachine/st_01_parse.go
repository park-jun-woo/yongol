//ff:func feature=validate type=rule control=iteration dimension=1 topic=statemachine-structural
//ff:what ST-1 — detects Mermaid stateDiagram parse errors by re-parsing and collecting diagnostics

package statemachine

import (
	"path/filepath"

	smparser "github.com/park-jun-woo/yongol/pkg/parser/statemachine"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// st01Parse re-parses states/ to surface diagnostics when fs.StateDiagrams
// came back empty or partial.
func st01Parse(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if len(fs.StateDiagrams) > 0 {
		return nil
	}
	_, diags := smparser.ParseDir(filepath.Join(fs.SpecsDir, "states"))
	for i := range diags {
		diags[i].Phase = diagnostic.PhaseValidate
		diags[i].Message = "[ST-1] " + diags[i].Message
		diags[i].Advice = "Check the Mermaid stateDiagram-v2 syntax"
	}
	return diags
}
