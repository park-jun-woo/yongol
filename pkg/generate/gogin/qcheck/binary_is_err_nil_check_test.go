//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestBinaryIsErrNilCheck — err!=nil / nil!=err / 비-ident / 무-err 분기 검증
package qcheck

import (
	"testing"
)

func TestBinaryIsErrNilCheck(t *testing.T) {
	cases := []struct {
		expr string
		want bool
	}{
		{"err != nil", true},
		{"nil != err", true},
		{"myErr != nil", true},
		{"x != nil", false},
		{"nil != x", false},
		{"nil != nil", false},
		{"err != other", false}, // neither side is nil
		{"foo() != nil", false}, // LHS not an Ident
	}
	for _, c := range cases {
		if got := binaryIsErrNilCheck(binExpr(t, c.expr)); got != c.want {
			t.Errorf("binaryIsErrNilCheck(%q) = %v, want %v", c.expr, got, c.want)
		}
	}
}
