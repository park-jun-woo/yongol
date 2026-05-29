//ff:func feature=validate-contract type=test-helper control=sequence topic=preserve-safety
//ff:what parse helpers — 테스트용 단일 expr/stmt/funcDecl AST 파싱 헬퍼

package contract

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

// mustExpr parses a standalone Go expression and returns its AST node.
func mustExpr(t *testing.T, src string) ast.Expr {
	t.Helper()
	e, err := parser.ParseExpr(src)
	if err != nil {
		t.Fatalf("parse expr %q: %v", src, err)
	}
	return e
}

// mustCall parses src as an expression and asserts it is a CallExpr.
func mustCall(t *testing.T, src string) *ast.CallExpr {
	t.Helper()
	call, ok := mustExpr(t, src).(*ast.CallExpr)
	if !ok {
		t.Fatalf("expr %q is not a CallExpr", src)
	}
	return call
}

// mustFirstStmt wraps body inside a func and returns the first stmt.
func mustFirstStmt(t *testing.T, body string) ast.Stmt {
	t.Helper()
	stmts := mustStmts(t, body)
	if len(stmts) == 0 {
		t.Fatalf("body %q has no statements", body)
	}
	return stmts[0]
}

// mustStmts wraps body inside a func declaration and returns all of
// its top-level statements.
func mustStmts(t *testing.T, body string) []ast.Stmt {
	t.Helper()
	return mustFuncDecl(t, "func _f() {\n"+body+"\n}").Body.List
}

// mustFuncDecl parses a single top-level func declaration source.
func mustFuncDecl(t *testing.T, src string) *ast.FuncDecl {
	t.Helper()
	full := "package p\n" + src + "\n"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", full, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse func %q: %v", src, err)
	}
	for _, d := range f.Decls {
		if fd, ok := d.(*ast.FuncDecl); ok {
			return fd
		}
	}
	t.Fatalf("no func decl in %q", src)
	return nil
}

// mustBlock parses body (a list of statements) wrapped in a func and
// returns the resulting *ast.BlockStmt.
func mustBlock(t *testing.T, body string) *ast.BlockStmt {
	t.Helper()
	return mustFuncDecl(t, "func _f() {\n"+body+"\n}").Body
}
