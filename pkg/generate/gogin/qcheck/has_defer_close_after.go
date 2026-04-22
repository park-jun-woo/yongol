//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what hasDeferCloseAfter — blockList[start+1:] 안에서 `defer <name>.Close()` 존재 여부 판정

package qcheck

import "go/ast"

// hasDeferCloseAfter walks the tail of blockList starting at start+1 and
// reports whether any statement is a `defer <name>.Close()` call on the
// given variable name. Only direct defer statements at block-level are
// accepted — defers nested inside anonymous func literals (a valid but
// rare lifetime-extension pattern) are ignored to keep the scanner simple
// and false-positive free for the codegen templates yongol emits.
func hasDeferCloseAfter(blockList []ast.Stmt, start int, name string) bool {
	for j := start + 1; j < len(blockList); j++ {
		def, ok := blockList[j].(*ast.DeferStmt)
		if !ok {
			continue
		}
		sel, ok := def.Call.Fun.(*ast.SelectorExpr)
		if !ok || sel.Sel.Name != "Close" {
			continue
		}
		recv, ok := sel.X.(*ast.Ident)
		if !ok {
			continue
		}
		if recv.Name == name {
			return true
		}
	}
	return false
}
