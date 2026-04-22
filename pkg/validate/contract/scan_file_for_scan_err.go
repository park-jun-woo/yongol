//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what scanFileForScanErr — preserved 파일에서 *sql.Row/Rows.Scan 에러 누락 진단 수집

package contract

import (
	"go/ast"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// scanFileForScanErr walks every BlockStmt in path and delegates per-
// block analysis to scanDiagsInBlock. We intentionally avoid parsing
// type information — receiver-type inference by name ("row", "rows",
// "r") keeps the rule precise enough for the generated handler shape
// without dragging in go/types (which would require the full module
// to be build-compilable).
func scanFileForScanErr(path string) []diagnostic.Diagnostic {
	fset, file, err := parseGoFile(path)
	if err != nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	ast.Inspect(file, func(n ast.Node) bool {
		block, ok := n.(*ast.BlockStmt)
		if !ok {
			return true
		}
		diags = append(diags, scanDiagsInBlock(fset, file, path, block)...)
		return true
	})
	return diags
}
