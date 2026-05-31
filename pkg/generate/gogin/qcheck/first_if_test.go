//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestIfInitCalls — if-init이 pkg.Func 호출인지 각 분기 검증
package qcheck

import (
	"go/ast"
	"testing"
)

func firstIf(t *testing.T, body string) *ast.IfStmt {
	t.Helper()
	for _, s := range blockStmts(t, body) {
		if ifs, ok := s.(*ast.IfStmt); ok {
			return ifs
		}
	}
	t.Fatalf("no if statement found in %q", body)
	return nil
}
