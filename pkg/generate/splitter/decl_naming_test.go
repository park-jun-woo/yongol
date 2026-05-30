//ff:func feature=gen-splitter type=test control=iteration dimension=1
//ff:what methodReceiver / receiverName / funcIdentifier / genDeclIdentifier / declIdentifier / funcFileName / genDeclFileName / fileNameForDecl

package splitter

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func declOf(t *testing.T, src string) ast.Decl {
	t.Helper()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	return f.Decls[0]
}

func TestMethodReceiverAndName(t *testing.T) {
	plain := declOf(t, "package p\nfunc Foo() {}").(*ast.FuncDecl)
	if got := methodReceiver(plain); got != "" {
		t.Errorf("plain func receiver = %q, want empty", got)
	}
	val := declOf(t, "package p\nfunc (q Queries) Bar() {}").(*ast.FuncDecl)
	if got := methodReceiver(val); got != "Queries" {
		t.Errorf("value receiver = %q, want Queries", got)
	}
	ptr := declOf(t, "package p\nfunc (q *Queries) Baz() {}").(*ast.FuncDecl)
	if got := methodReceiver(ptr); got != "Queries" {
		t.Errorf("pointer receiver = %q, want Queries", got)
	}
}

func TestReceiverNameExotic(t *testing.T) {
	if got := receiverName(&ast.IndexExpr{X: &ast.Ident{Name: "G"}}); got != "" {
		t.Errorf("generic receiver = %q, want empty", got)
	}
}

func TestFuncIdentifier(t *testing.T) {
	plain := declOf(t, "package p\nfunc Foo() {}").(*ast.FuncDecl)
	if got := funcIdentifier(plain); got != "Foo" {
		t.Errorf("got %q, want Foo", got)
	}
	method := declOf(t, "package p\nfunc (q *Queries) FindUser() {}").(*ast.FuncDecl)
	if got := funcIdentifier(method); got != "Queries.FindUser" {
		t.Errorf("got %q, want Queries.FindUser", got)
	}
}

func TestGenDeclIdentifier(t *testing.T) {
	typ := declOf(t, "package p\ntype Row struct{}").(*ast.GenDecl)
	if got := genDeclIdentifier(typ); got != "Row" {
		t.Errorf("got %q, want Row", got)
	}
	multi := declOf(t, "package p\nconst (\nA=1\nB=2\n)").(*ast.GenDecl)
	if got := genDeclIdentifier(multi); got != "const" {
		t.Errorf("got %q, want const", got)
	}
}

func TestDeclIdentifier(t *testing.T) {
	fn := declOf(t, "package p\nfunc Foo() {}")
	if got := declIdentifier(fn); got != "Foo" {
		t.Errorf("func got %q", got)
	}
	gd := declOf(t, "package p\ntype Row struct{}")
	if got := declIdentifier(gd); got != "Row" {
		t.Errorf("gendecl got %q", got)
	}
}

func TestFuncFileName(t *testing.T) {
	plain := declOf(t, "package p\nfunc FindUser() {}").(*ast.FuncDecl)
	if got := funcFileName(plain, ".go"); got != "find_user.go" {
		t.Errorf("plain = %q, want find_user.go", got)
	}
	method := declOf(t, "package p\nfunc (q *Queries) FindUser() {}").(*ast.FuncDecl)
	if got := funcFileName(method, ".sql.go"); got != "queries_find_user.sql.go" {
		t.Errorf("method = %q, want queries_find_user.sql.go", got)
	}
}

func TestGenDeclFileName(t *testing.T) {
	typ := declOf(t, "package p\ntype Row struct{}").(*ast.GenDecl)
	if got := genDeclFileName(typ, ".go"); got != "row.go" {
		t.Errorf("type = %q, want row.go", got)
	}
	constBlock := declOf(t, "package p\nconst (\nA=1\nB=2\n)").(*ast.GenDecl)
	if got := genDeclFileName(constBlock, ".go"); got != "consts.go" {
		t.Errorf("const block = %q, want consts.go", got)
	}
	varBlock := declOf(t, "package p\nvar (\nA int\nB int\n)").(*ast.GenDecl)
	if got := genDeclFileName(varBlock, ".go"); got != "vars.go" {
		t.Errorf("var block = %q, want vars.go", got)
	}
}

func TestFileNameForDecl(t *testing.T) {
	fn := declOf(t, "package p\nfunc FindUser() {}")
	if got := fileNameForDecl(fn, ToolSQLC, false); got != "find_user.sql.go" {
		t.Errorf("func = %q, want find_user.sql.go", got)
	}
	typ := declOf(t, "package p\ntype Row struct{}")
	if got := fileNameForDecl(typ, ToolSQLC, true); got != "row.model.go" {
		t.Errorf("models type = %q, want row.model.go", got)
	}
	if got := fileNameForDecl(fn, ToolOAPICodegen, false); got != "find_user.gen.go" {
		t.Errorf("oapi func = %q, want find_user.gen.go", got)
	}
}
