//ff:func feature=gen-splitter type=test control=iteration dimension=1
//ff:what docOf / funcDoc / genDeclDoc / detectControl / controlFor / funcTypeFor / extractHeader
package splitter

import (
	"go/ast"
	"testing"
)

func TestDetectControl(t *testing.T) {
	cases := []struct {
		name    string
		body    string
		wantC   string
		wantDim string
	}{
		{"sequence", "x := 1\n_ = x", "sequence", ""},
		{"loop", "for i := 0; i < 3; i++ {}", "iteration", "1"},
		{"range", "for range []int{} {}", "iteration", "1"},
		{"switch", "switch {\ncase true:\n}", "selection", ""},
		{"both falls to sequence", "for {}\nswitch {\n}", "sequence", ""},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			fn := parseDeclWithComments(t, "package p\nfunc f() {\n"+c.body+"\n}").(*ast.FuncDecl)
			gotC, gotDim := detectControl(fn.Body)
			if gotC != c.wantC || gotDim != c.wantDim {
				t.Errorf("detectControl = (%q,%q), want (%q,%q)", gotC, gotDim, c.wantC, c.wantDim)
			}
		})
	}
	// nil body
	if c, d := detectControl(nil); c != "sequence" || d != "" {
		t.Errorf("nil body = (%q,%q)", c, d)
	}
}
