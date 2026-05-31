//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what parse helpers — 테스트용 단일 expr/stmt/funcDecl AST 파싱 헬퍼
package contract

import (
	"go/ast"
	"testing"
)

func mustStmts(t *testing.T, body string) []ast.Stmt {
	t.Helper()
	return mustFuncDecl(t, "func _f() {\n"+body+"\n}").Body.List
}
