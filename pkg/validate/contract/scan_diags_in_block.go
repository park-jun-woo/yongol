//ff:func feature=validate-contract type=util control=iteration dimension=1 topic=preserve-safety
//ff:what scanDiagsInBlock — 단일 블록 내 Scan 호출 에러 누락 진단 생성

package contract

import (
	"go/ast"
	"go/token"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// scanDiagsInBlock iterates stmts and raises PRV-13 for each Scan
// call whose returned error is dropped. Recognised safe forms mirror
// PRV-12: explicit discard, inline `if err := ...Scan(...); err != nil`,
// or a later `if err != nil` guard on the same errName. Per-stmt
// analysis is delegated to scanStmtDiag so this orchestrator stays
// short enough for Q4.
func scanDiagsInBlock(fset *token.FileSet, file *ast.File, path string, block *ast.BlockStmt) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for idx, stmt := range block.List {
		if d, ok := scanStmtDiag(fset, file, path, block, idx, stmt); ok {
			diags = append(diags, d)
		}
	}
	return diags
}
