//ff:func feature=funcspec type=parser control=iteration dimension=1
//ff:what extractReturnTypes — FuncDecl 의 반환 타입 위치별로 펼쳐 문자열 슬라이스로

package funcspec

import (
	"go/ast"
	"go/printer"
	"go/token"
	"strings"
)

// extractReturnTypes prints each result type from the FuncDecl's signature so
// that @eval (S-67) can verify a predicate func's return shape is `bool`.
// Named-results are expanded one entry per name; anonymous results contribute
// one entry. Returns nil when the function declares no results.
func extractReturnTypes(fset *token.FileSet, d *ast.FuncDecl) []string {
	if d.Type == nil || d.Type.Results == nil {
		return nil
	}
	var out []string
	for _, field := range d.Type.Results.List {
		var buf strings.Builder
		_ = printer.Fprint(&buf, fset, field.Type)
		typ := buf.String()
		count := len(field.Names)
		if count == 0 {
			count = 1
		}
		out = appendN(out, typ, count)
	}
	return out
}
