//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what parse helpers — 테스트용 단일 expr/stmt/funcDecl AST 파싱 헬퍼
package contract

import (
	"go/ast"
	"go/parser"
	"testing"
)

func mustExpr(t *testing.T, src string) ast.Expr {
	t.Helper()
	e, err := parser.ParseExpr(src)
	if err != nil {
		t.Fatalf("parse expr %q: %v", src, err)
	}
	return e
}
