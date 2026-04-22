//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-25 — unknown sequence type

package ssac

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s25UnknownSeqType validates S-25: every seq.Type must be a known directive.
func s25UnknownSeqType(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if knownSeqTypes[seq.Type] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-25] unknown sequence type: @%s", seq.Type),
				Advice:  "Use one of: @get/@post/@put/@delete/@call/@empty/@exists/@state/@auth/@publish/@verify-password",
			})
		}
	}
	return diags
}
