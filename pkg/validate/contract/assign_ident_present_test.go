//ff:func feature=validate-contract type=test-helper control=iteration dimension=1 topic=preserve-safety
//ff:what assignIdentPresent — flatten 결과에 단일-LHS 이름 name 의 AssignStmt 가 있는지 검사 헬퍼
package contract

import "go/ast"

// assignIdentPresent reports whether stmts contains an AssignStmt whose single
// LHS is an identifier named name.
func assignIdentPresent(stmts []ast.Stmt, name string) bool {
	for _, s := range stmts {
		as, ok := s.(*ast.AssignStmt)
		if !ok || len(as.Lhs) != 1 {
			continue
		}
		if id, ok := as.Lhs[0].(*ast.Ident); ok && id.Name == name {
			return true
		}
	}
	return false
}
