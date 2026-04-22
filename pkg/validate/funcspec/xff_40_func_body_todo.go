//ff:func feature=validate type=rule control=iteration dimension=1 topic=funcspec-structural
//ff:what XFF-40 — detects funcs whose body is not yet implemented (HasBody=false)

package funcspec

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xff40FuncBodyTodo flags FuncSpecs whose parser marked HasBody=false
// (panic("TODO") / // TODO / empty / simple zero-return).
func xff40FuncBodyTodo(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, sp := range fs.ProjectFuncSpecs {
		if sp.HasBody {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    sp.Package + "/" + sp.Name + ".go",
			Line:    sp.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[XFF-40] func " + sp.Package + "." + sp.Name + " is a stub (unimplemented)",
			Advice:  "Remove the TODO comment and implement the function body",
		})
	}
	return diags
}
