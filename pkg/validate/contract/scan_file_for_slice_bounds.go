//ff:func feature=validate-contract type=util control=iteration dimension=1 topic=preserve-safety
//ff:what scanFileForSliceBounds — preserved 파일 FuncDecl 단위로 slice[0] 가드 누락 수집

package contract

import (
	"go/ast"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// scanFileForSliceBounds iterates function declarations and, for each
// body, hands off to sliceBoundsDiagsInFunc so "observed guards" are
// scoped per function rather than per file — this is what keeps the
// rule precise for handler files that declare multiple functions.
func scanFileForSliceBounds(path string) []diagnostic.Diagnostic {
	fset, file, err := parseGoFile(path)
	if err != nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Body == nil {
			continue
		}
		diags = append(diags, sliceBoundsDiagsInFunc(fset, file, path, fn.Body)...)
	}
	return diags
}
