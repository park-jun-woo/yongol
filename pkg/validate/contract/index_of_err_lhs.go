//ff:func feature=validate-contract type=util control=iteration dimension=1 topic=preserve-safety
//ff:what indexOfErrLhs — AssignStmt LHS 중 errName 에 해당하는 ident 위치 반환

package contract

import "go/ast"

// indexOfErrLhs returns the 0-based LHS position where errName (or
// "err" when errName is empty) is the identifier. A return of -1
// means the assignment does not touch the tracked error variable, so
// assignClobbersErr can short-circuit to "no clobber".
func indexOfErrLhs(assign *ast.AssignStmt, errName string) int {
	for i, lhs := range assign.Lhs {
		ident, ok := lhs.(*ast.Ident)
		if !ok {
			continue
		}
		if errName != "" && ident.Name == errName {
			return i
		}
		if errName == "" && ident.Name == "err" {
			return i
		}
	}
	return -1
}
