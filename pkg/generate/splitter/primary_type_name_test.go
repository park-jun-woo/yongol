//ff:func feature=gen-splitter type=test control=iteration dimension=1
//ff:what primaryTypeName / sqlcSuffix / suffixFor / matchesOriginal / importName AST·도구 헬퍼
package splitter

import (
	"go/ast"
	"testing"
)

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
