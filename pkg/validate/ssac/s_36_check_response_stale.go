//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what s36CheckResponseStale — detects stale variable references in @response fields (WARNING)

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// s36CheckResponseStale emits WARNING diagnostics for @response sequences
// that reference variables currently marked stale (mutated but not re-queried).
func s36CheckResponseStale(fn parsessac.ServiceFunc, _ int, seq parsessac.Sequence, stale map[string]bool) []diagnostic.Diagnostic {
	if seq.Type != "response" || seq.SuppressWarn {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, varRef := range seq.Fields {
		ref := inputValueRefBase(varRef)
		if ref == "" || !stale[ref] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    fn.FileName,
			Line:    seq.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: fmt.Sprintf("[S-36] @response uses %s which was mutated but not re-queried", ref),
			Advice:  "Re-fetch the modified object with @get after a @put/@delete",
		})
	}
	return diags
}
