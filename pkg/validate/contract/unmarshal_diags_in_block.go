//ff:func feature=validate-contract type=util control=iteration dimension=1 topic=preserve-safety
//ff:what unmarshalDiagsInBlock — 단일 블록 내 Unmarshal 에러 미처리 진단 생성

package contract

import (
	"go/ast"
	"go/token"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// unmarshalDiagsInBlock walks block.List and emits PRV-12 diagnostics
// for Unmarshal-style statements that silently discard the returned
// error. Accepted safe shapes:
//
//   - `err := json.Unmarshal(...)` followed later in the block by an
//     `if err != nil { ... }` guard.
//   - `if err := json.Unmarshal(...); err != nil { ... }` — err is
//     scoped to the if init and checked inline.
//   - `_ = json.Unmarshal(...)` — explicit discard, documented intent.
//   - `_, _ = ...` — same.
func unmarshalDiagsInBlock(fset *token.FileSet, file *ast.File, path string, block *ast.BlockStmt) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for idx, stmt := range block.List {
		_, call, errName, kind := unmarshalAssignInStmt(stmt)
		if call == nil || kind == unmarshalKindDiscarded {
			continue
		}
		line := fset.Position(call.Pos()).Line
		if hasNolint(fset, file, line, "prv-12") {
			continue
		}
		if errName != "" && hasErrCheckAfter(block.List, idx, errName) {
			continue
		}
		diags = append(diags, makeUnmarshalDiag(path, line))
	}
	return diags
}
