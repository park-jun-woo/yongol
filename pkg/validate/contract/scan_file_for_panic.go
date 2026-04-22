//ff:func feature=validate-contract type=util control=iteration dimension=1 topic=preserve-safety
//ff:what scanFileForPanic — 단일 preserved 파일에서 허용되지 않은 panic() 호출 수집

package contract

import (
	"go/ast"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// scanFileForPanic parses path once and walks every non-init func body
// for `panic(...)` calls that lack a `// nolint:panic` escape comment.
// Parse errors are swallowed — contract PRV-01/02 already surface real
// syntax failures, so PRV-10 stays silent to avoid double reporting.
func scanFileForPanic(path string) []diagnostic.Diagnostic {
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
		if isInitFunc(fn) {
			continue
		}
		diags = append(diags, panicDiagsInBody(fset, file, path, fn.Body)...)
	}
	return diags
}
