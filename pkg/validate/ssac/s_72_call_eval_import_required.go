//ff:func feature=validate type=rule control=iteration dimension=2 topic=ssac-structural
//ff:what S-72 — @call/@eval 참조 패키지에 대한 SSaC import 선언 필수

package ssac

import (
	"fmt"
	"path"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// s72CallEvalImportRequired validates S-72: every @call and @eval sequence
// that references a dotted model (pkg.Func) must have a matching import
// declaration in the enclosing ServiceFunc. A match is defined as
// path.Base(importPath) == pkgAlias.
func s72CallEvalImportRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		// Build a set of imported package aliases for fast lookup.
		importedAliases := make(map[string]bool, len(fn.Imports))
		for _, imp := range fn.Imports {
			importedAliases[path.Base(imp)] = true
		}
		for _, seq := range fn.Sequences {
			if seq.Type != parsessac.SeqCall && seq.Type != parsessac.SeqEval {
				continue
			}
			if seq.Model == "" {
				continue
			}
			dotIdx := strings.IndexByte(seq.Model, '.')
			if dotIdx < 0 {
				continue // no package qualifier — other rules handle this
			}
			pkgAlias := seq.Model[:dotIdx]
			if importedAliases[pkgAlias] {
				continue
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:    fn.FileName,
				Line:    seq.Line,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[S-72] @%s references package %q but no matching import declaration found", seq.Type, pkgAlias),
				Advice:  fmt.Sprintf("Add import %q with the full Go import path to the top of the SSaC file.", pkgAlias),
			})
		}
	}
	return diags
}
