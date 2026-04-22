//ff:func feature=validate type=rule control=iteration dimension=3 topic=ssac-structural
//ff:what S-37 — a FK-referenced @get result requires an @empty guard

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s37FKReferenceGuard validates S-37: a single-row @get that takes a foreign-key
// reference to a previously declared variable should be followed by an @empty guard.
func s37FKReferenceGuard(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		declared := map[string]bool{}
		types := map[string]string{}
		if fn.Subscribe != nil {
			declared["message"] = true
		}
		for i, seq := range fn.Sequences {
			if seq.Type == "get" && seq.Result != nil &&
				!strings.HasPrefix(seq.Result.Type, "[]") && seq.Result.Wrapper == "" {
				model := extractModel(seq)
				if hasFKRef(seq, declared, types, model) && !hasEmptyGuardAfter(fn.Sequences[i+1:], seq.Result.Var) {
					diags = append(diags, diagnostic.Diagnostic{
						File:    fn.FileName,
						Line:    seq.Line,
						Phase:   diagnostic.PhaseValidate,
						Level:   diagnostic.LevelWarning,
						Message: fmt.Sprintf("[S-37] %s: FK reference @get requires @empty guard", seq.Result.Var),
						Advice:  "Add an @empty guard for the object fetched with @get",
					})
				}
			}
			if seq.Result != nil {
				declared[seq.Result.Var] = true
				types[seq.Result.Var] = seq.Result.Type
			}
		}
	}
	return diags
}
