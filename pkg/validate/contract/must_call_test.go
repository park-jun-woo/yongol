//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what parse helpers — 테스트용 단일 expr/stmt/funcDecl AST 파싱 헬퍼
package contract

import (
	"go/ast"
	"testing"
)

func mustCall(t *testing.T, src string) *ast.CallExpr {
	t.Helper()
	call, ok := mustExpr(t, src).(*ast.CallExpr)
	if !ok {
		t.Fatalf("expr %q is not a CallExpr", src)
	}
	return call
}
