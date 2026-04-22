//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what panicDiagsInBody — 함수 body AST 를 walk 하여 panic() 호출마다 PRV-10 Diagnostic 생성

package contract

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// panicDiagsInBody walks body recursively and emits one PRV-10
// Diagnostic per unsuppressed `panic(...)` CallExpr. Suppression works
// via `// nolint:panic` on the call's own line or the line directly
// above it — hasNolint handles both.
func panicDiagsInBody(fset *token.FileSet, file *ast.File, path string, body *ast.BlockStmt) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	ast.Inspect(body, func(n ast.Node) bool {
		call, ok := n.(*ast.CallExpr)
		if !ok {
			return true
		}
		ident, ok := call.Fun.(*ast.Ident)
		if !ok || ident.Name != "panic" {
			return true
		}
		line := fset.Position(call.Pos()).Line
		if hasNolint(fset, file, line, "panic") {
			return true
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    path,
			Line:    line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[PRV-10] preserved file introduces panic() in production path (line %d)", line),
			Advice: "Original generated code returned an error; return one here instead.\n" +
				"If intentional, add `// nolint:panic` on the line above.",
		})
		return true
	})
	return diags
}
