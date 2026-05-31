//ff:func feature=funcspec type=test control=iteration dimension=1
//ff:what extractReturnTypes / isStubBody / processFuncDecl / findFuncDeclLine / extractGoParseErrorLine
package funcspec

import (
	"testing"
)

func TestExtractReturnTypes(t *testing.T) {
	cases := []struct {
		name string
		src  string
		want []string
	}{
		{"no results", "package p\nfunc f() {}", nil},
		{"single", "package p\nfunc f() error { return nil }", []string{"error"}},
		{"two anon", "package p\nfunc f() (Resp, error) { return Resp{}, nil }", []string{"Resp", "error"}},
		{"named group expands", "package p\nfunc f() (a, b int) { return }", []string{"int", "int"}},
		{"pointer", "package p\nfunc f() *T { return nil }", []string{"*T"}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			assertExtractReturnTypes(t, c.src, c.want)
		})
	}
}
