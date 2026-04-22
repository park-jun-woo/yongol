//ff:func feature=validate type=util control=iteration dimension=1 topic=funcspec-structural
//ff:what collectForbiddenImports — generates XFF-41 diagnostics for forbidden imports in a single FuncSpec

package funcspec

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

func collectForbiddenImports(pkg, name string, line int, imports []string) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, imp := range imports {
		if !isForbiddenImport(imp) {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    pkg + "/" + name + ".go",
			Line:    line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[XFF-41] func " + pkg + "." + name + " imports forbidden I/O package " + imp,
			Advice:  "Remove the forbidden import and use the yongol-provided pkg/<X> instead",
		})
	}
	return diags
}
