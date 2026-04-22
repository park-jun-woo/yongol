//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what scanStmtDiag — 단일 stmt 가 PRV-13 위반이면 Diagnostic 생성

package contract

import (
	"go/ast"
	"go/token"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// scanStmtDiag inspects stmt at index idx within block.List and
// returns (diag, true) when it represents an unguarded Scan call.
// All the fast-path skip conditions live here: not-a-scan, explicitly
// discarded, covered by a nolint comment, or guarded by a later
// `if err != nil` in the same block.
func scanStmtDiag(fset *token.FileSet, file *ast.File, path string, block *ast.BlockStmt, idx int, stmt ast.Stmt) (diagnostic.Diagnostic, bool) {
	call, errName, kind := scanCallInStmt(stmt)
	if call == nil || kind == scanKindDiscarded {
		return diagnostic.Diagnostic{}, false
	}
	line := fset.Position(call.Pos()).Line
	if hasNolint(fset, file, line, "prv-13") {
		return diagnostic.Diagnostic{}, false
	}
	if errName != "" && hasErrCheckAfter(block.List, idx, errName) {
		return diagnostic.Diagnostic{}, false
	}
	return makeScanDiag(path, line), true
}
