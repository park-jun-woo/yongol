//ff:func feature=validate-contract type=util control=iteration dimension=1 topic=preserve-safety
//ff:what scanFileForMissingClose — preserved 파일 함수별 리소스 획득/close 매칭 검증

package contract

import (
	"go/ast"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// scanFileForMissingClose iterates FuncDecls so the close-coverage
// analysis is scoped to a single function body — exactly like
// sliceBoundsDiagsInFunc. Nested functions (FuncLits) are handled
// eagerly by the walker inside missingCloseDiagsInFunc because a
// FuncLit body is its own resource lifecycle.
func scanFileForMissingClose(path string) []diagnostic.Diagnostic {
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
		diags = append(diags, missingCloseDiagsInFunc(fset, file, path, fn.Body)...)
	}
	return diags
}
