//ff:func feature=validate-contract type=util control=selection topic=preserve-safety
//ff:what isScanCall — call 이 `<row|rows|r>.Scan(...)` 호출인지 판정

package contract

import (
	"go/ast"
	"strings"
)

// isScanCall recognises `.Scan(...)` on a receiver whose identifier
// looks like a sql.Row / sql.Rows handle. Accepted receiver names:
// "row", "rows", "r" (sqlc-generated shape), or any identifier ending
// in "Row"/"Rows"/"row"/"rows". This avoids false positives on
// unrelated methods named Scan (e.g. bufio.Scanner.Scan which is
// parameter-less and returns bool, not error).
func isScanCall(call *ast.CallExpr) bool {
	if call == nil || len(call.Args) == 0 {
		return false
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel == nil || sel.Sel.Name != "Scan" {
		return false
	}
	recvName := leftmostIdentName(sel.X)
	if recvName == "" {
		return false
	}
	switch recvName {
	case "row", "rows", "r":
		return true
	}
	lower := strings.ToLower(recvName)
	return strings.HasSuffix(lower, "row") || strings.HasSuffix(lower, "rows")
}
