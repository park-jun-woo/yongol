//ff:func feature=gen-splitter type=test control=iteration dimension=1
//ff:what primaryTypeName / sqlcSuffix / suffixFor / matchesOriginal / importName AST·도구 헬퍼

package splitter

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func firstDecl(t *testing.T, src string) ast.Decl {
	t.Helper()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	return f.Decls[0]
}

func TestPrimaryTypeName(t *testing.T) {
	cases := []struct {
		name string
		src  string
		want string
	}{
		{"single type", "package p\ntype Foo struct{}", "Foo"},
		{"single const", "package p\nconst X = 1", "X"},
		{"single var", "package p\nvar Y int", "Y"},
		{"multi spec", "package p\nconst (\nA = 1\nB = 2\n)", ""},
		{"multi name value", "package p\nvar a, b int", ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gd := firstDecl(t, c.src).(*ast.GenDecl)
			if got := primaryTypeName(gd); got != c.want {
				t.Errorf("primaryTypeName = %q, want %q", got, c.want)
			}
		})
	}
}

func TestSqlcSuffix(t *testing.T) {
	typeDecl := firstDecl(t, "package p\ntype Row struct{}")
	funcDecl := firstDecl(t, "package p\nfunc f() {}")
	if got := sqlcSuffix(false, typeDecl); got != ".sql.go" {
		t.Errorf("non-models type = %q, want .sql.go", got)
	}
	if got := sqlcSuffix(true, typeDecl); got != ".model.go" {
		t.Errorf("models type = %q, want .model.go", got)
	}
	if got := sqlcSuffix(true, funcDecl); got != ".sql.go" {
		t.Errorf("models func = %q, want .sql.go", got)
	}
}

func TestSuffixFor(t *testing.T) {
	typeDecl := firstDecl(t, "package p\ntype Row struct{}")
	if got := suffixFor(ToolOAPICodegen, false, typeDecl); got != ".gen.go" {
		t.Errorf("oapi = %q, want .gen.go", got)
	}
	if got := suffixFor(ToolSQLC, true, typeDecl); got != ".model.go" {
		t.Errorf("sqlc models = %q, want .model.go", got)
	}
	if got := suffixFor(Tool("unknown"), false, typeDecl); got != ".go" {
		t.Errorf("default = %q, want .go", got)
	}
}

func TestMatchesOriginal(t *testing.T) {
	cases := []struct {
		name string
		tool Tool
		want bool
	}{
		{"api.gen.go", ToolOAPICodegen, true},
		{"api.go", ToolOAPICodegen, false},
		{"models.go", ToolSQLC, true},
		{"users.sql.go", ToolSQLC, true},
		{"users.model.go", ToolSQLC, true},
		{"querier.go", ToolSQLC, false},
		{"x.go", Tool("other"), false},
	}
	for _, c := range cases {
		if got := matchesOriginal(c.name, c.tool); got != c.want {
			t.Errorf("matchesOriginal(%q,%v) = %v, want %v", c.name, c.tool, got, c.want)
		}
	}
}

func importSpecOf(t *testing.T, src string) *ast.ImportSpec {
	t.Helper()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	gd := f.Decls[0].(*ast.GenDecl)
	return gd.Specs[0].(*ast.ImportSpec)
}

func TestImportName(t *testing.T) {
	cases := []struct {
		name string
		src  string
		want string
	}{
		{"plain", "package p\nimport \"fmt\"", "fmt"},
		{"path base", "package p\nimport \"net/http\"", "http"},
		{"alias", "package p\nimport x \"net/http\"", "x"},
		{"version suffix", "package p\nimport \"github.com/foo/bar/v2\"", "bar"},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			if got := importName(importSpecOf(t, c.src)); got != c.want {
				t.Errorf("importName = %q, want %q", got, c.want)
			}
		})
	}
}
