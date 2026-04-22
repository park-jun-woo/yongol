//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-28 — the Target variable in @empty must be declared

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s28TargetDeclared validates S-28: Target's leading variable must be declared.
func s28TargetDeclared(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		for i, seq := range fn.Sequences {
			if seq.Target == "" {
				continue
			}
			ref := strings.SplitN(seq.Target, ".", 2)[0]
			if ref == "" || isImplicitVar(ref) {
				continue
			}
			declared := declaredVars(fn, i)
			if declared[ref] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-28] Target variable %q used before declaration", ref),
				Advice:  fmt.Sprintf("Declare variable %q as the result of a preceding @get/@post sequence", ref),
			})
		}
	}
	return diags
}
