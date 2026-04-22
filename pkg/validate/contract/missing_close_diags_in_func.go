//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what missingCloseDiagsInFunc — 함수 body 에서 획득-then-close 매칭 검증, 누락 시 PRV-17

package contract

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// missingCloseDiagsInFunc walks body once and for every assignment of
// the form `v, err := <resource-call>(...)` (or single-LHS variant)
// emits PRV-17 unless the same body carries a `defer v.Close()`. The
// body's full statement list is captured up front so hasDeferClose
// can scan forward or backward — defers may appear before or after
// the acquisition as long as they are in the same function body.
func missingCloseDiagsInFunc(fset *token.FileSet, file *ast.File, path string, body *ast.BlockStmt) []diagnostic.Diagnostic {
	stmts := flattenBlockStmts(body)
	var diags []diagnostic.Diagnostic
	ast.Inspect(body, func(n ast.Node) bool {
		as, ok := n.(*ast.AssignStmt)
		if !ok {
			return true
		}
		varName, call := resourceAcquireFromAssign(as)
		if call == nil || varName == "" {
			return true
		}
		line := fset.Position(call.Pos()).Line
		if hasNolint(fset, file, line, "prv-17") {
			return true
		}
		if hasDeferClose(stmts, varName) {
			return true
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    path,
			Line:    line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[PRV-17] preserved file acquires %s without a matching defer %s.Close() (line %d)", varName, varName, line),
			Advice: fmt.Sprintf("Close the resource before returning:\n"+
				"  %s, err := ...\n"+
				"  if err != nil { return api.Error500, err }\n"+
				"  defer %s.Close()", varName, varName),
		})
		return true
	})
	return diags
}
