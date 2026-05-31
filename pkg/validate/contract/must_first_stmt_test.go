//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what parse helpers — 테스트용 단일 expr/stmt/funcDecl AST 파싱 헬퍼
package contract

import (
	"go/ast"
	"testing"
)

func mustFirstStmt(t *testing.T, body string) ast.Stmt {
	t.Helper()
	stmts := mustStmts(t, body)
	if len(stmts) == 0 {
		t.Fatalf("body %q has no statements", body)
	}
	return stmts[0]
}
