//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what XSS-11 — @result type is in plural form (WARNING)

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xss11PluralResultType validates XSS-11: when the wrapper (Page/Cursor/[])
// already conveys plurality, the inner element type T must be singular.
// `Page[Gig]` and `[]Gig` are correct; `Page[Gigs]` and `[]Gigs` are
// suspicious. Likewise a bare singular result whose type name is plural
// (e.g. Result.Type = "Gigs" with no wrapper) is flagged.
//
// Skip cases: @call seq, package-prefixed Model, primitive types.
func xss11PluralResultType(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	primitives := map[string]bool{
		"string": true, "int": true, "int32": true, "int64": true,
		"float32": true, "float64": true, "bool": true, "byte": true, "rune": true,
	}
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for _, seq := range fn.Sequences {
			if seq.Type == "call" {
				continue
			}
			if seq.Package != "" {
				continue
			}
			if seq.Result == nil || seq.Result.Type == "" {
				continue
			}
			t := stripTypePrefix(seq.Result.Type)
			if primitives[strings.ToLower(t)] {
				continue
			}
			endsInS := strings.HasSuffix(t, "s") || strings.HasSuffix(t, "S")
			if !endsInS {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[XSS-11] result element type %q is plural; element types should be singular (use Page[T]/Cursor[T]/[]T to convey plurality)", t),
				Advice:  "Change to the singular type name (e.g. Gig)",
			})
		}
	}
	return diags
}
