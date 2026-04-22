//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what assignLHSHasErr — AssignStmt LHS 에 err 계열 이름의 Ident 가 있는지 판정

package qcheck

import (
	"go/ast"
	"strings"
)

// assignLHSHasErr reports whether any LHS entry of assign is a plain
// identifier whose lowercased name equals "err" or ends with "err". Used
// to distinguish a genuine error capture (`err := foo()`) from a discard
// (`_ = foo()`) so the latter falls through to the "unchecked" branch.
func assignLHSHasErr(assign *ast.AssignStmt) bool {
	for _, lhs := range assign.Lhs {
		ident, ok := lhs.(*ast.Ident)
		if !ok {
			continue
		}
		lower := strings.ToLower(ident.Name)
		if lower == "err" || strings.HasSuffix(lower, "err") {
			return true
		}
	}
	return false
}
