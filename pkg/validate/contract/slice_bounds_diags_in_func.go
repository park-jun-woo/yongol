//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what sliceBoundsDiagsInFunc — 함수 body 내 가드 없는 `x[0]` 접근 진단 생성

package contract

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// sliceBoundsDiagsInFunc walks body once, recording names seen in
// `len(x)` positions (any comparison operand) and names seen inside
// `for _, v := range x { ... }` headers — both are evidence that the
// author has acknowledged non-emptiness before a subsequent `x[0]`
// access. Names inside a range loop's body are safe from index 0 for
// the iteration variable, not the range target — we track the target.
//
// We flag `x[0]` IndexExpr only when the indexed identifier is a
// plain *ast.Ident and has not been guarded earlier. This yields
// false-negative on multi-return expressions like `foo()[0]` which we
// intentionally skip (too many legitimate uses like string literal
// indexing).
func sliceBoundsDiagsInFunc(fset *token.FileSet, file *ast.File, path string, body *ast.BlockStmt) []diagnostic.Diagnostic {
	guarded := map[string]bool{}
	var diags []diagnostic.Diagnostic
	collectGuards(body, guarded)
	ast.Inspect(body, func(n ast.Node) bool {
		idx, ok := n.(*ast.IndexExpr)
		if !ok {
			return true
		}
		name := leftmostIdentName(idx.X)
		if name == "" || guarded[name] {
			return true
		}
		if !isZeroIndex(idx.Index) {
			return true
		}
		line := fset.Position(idx.Pos()).Line
		if hasNolint(fset, file, line, "prv-14") {
			return true
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    path,
			Line:    line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[PRV-14] preserved file accesses %s[0] without len guard (line %d)", name, line),
			Advice: fmt.Sprintf("Guard before indexing:\n"+
				"  if len(%s) == 0 { return api.Error404, nil }\n"+
				"  first := %s[0]", name, name),
		})
		return true
	})
	return diags
}
