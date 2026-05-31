//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestAssignLHSHasErr — err/접미사 err / 비-ident / 무-err LHS 분기 검증
package qcheck

import (
	"go/ast"
	"testing"
)

func TestAssignLHSHasErr(t *testing.T) {
	cases := []struct {
		name string
		body string
		want bool
	}{
		{"PlainErr", "err := f()", true},
		{"SuffixErr", "myErr := f()", true},
		{"DiscardUnderscore", "_ = f()", false},
		{"NoErrName", "x := f()", false},
		{"SelectorLHS", "obj.field = f()", false},
		{"MultiWithErr", "x, err := g()", true},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			list := blockStmts(t, c.body)
			assign, ok := list[0].(*ast.AssignStmt)
			if !ok {
				t.Fatalf("stmt is not AssignStmt: %T", list[0])
			}
			if got := assignLHSHasErr(assign); got != c.want {
				t.Errorf("assignLHSHasErr(%q) = %v, want %v", c.body, got, c.want)
			}
		})
	}
}
