//ff:func feature=validate type=rule control=iteration dimension=1 topic=funcspec-structural
//ff:what XFF-41 — func 파일이 I/O 경계 패키지를 import 하는지 감지

package funcspec

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xff41FuncForbiddenImport flags FuncSpecs that import forbidden I/O packages.
func xff41FuncForbiddenImport(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, sp := range fs.ProjectFuncSpecs {
		diags = append(diags, collectForbiddenImports(sp.Package, sp.Name, sp.Line, sp.Imports)...)
	}
	return diags
}
