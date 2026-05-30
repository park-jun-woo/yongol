//ff:func feature=gen-gogin type=test control=branch topic=depth-report
//ff:what TestMaxBlockDepth — for/range/switch/typeswitch/select/case 각 노드 종류별 depth 검증

package qcheck

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

// bodyBlock parses func H's body BlockStmt.
func bodyBlock(t *testing.T, body string) *ast.BlockStmt {
	t.Helper()
	src := "package x\nfunc H() {\n" + body + "\n}"
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	return file.Decls[0].(*ast.FuncDecl).Body
}

func TestMaxBlockDepth_Kinds(t *testing.T) {
	cases := []struct {
		name string
		body string
		want int
	}{
		{"flat", "_ = 1", 0},
		{"for", "for i := 0; i < 1; i++ { _ = i }", 1},
		{"range", "for _, v := range []int{1} { _ = v }", 1},
		{"switch", "switch { case true: _ = 1 }", 1},
		{"typeswitch", "var a any; switch a.(type) { case int: _ = 1 }", 1},
		{"select", "select { default: _ = 1 }", 1},
		{"nestedForInIf", "if c() { for { break } }", 2},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := maxBlockDepth(bodyBlock(t, c.body), 0)
			if got != c.want {
				t.Errorf("maxBlockDepth(%q) = %d, want %d", c.body, got, c.want)
			}
		})
	}
}

func TestMaxBlockDepth_NonStmtNode(t *testing.T) {
	// A non-control node returns the input depth unchanged.
	if got := maxBlockDepth(&ast.Ident{Name: "x"}, 7); got != 7 {
		t.Errorf("maxBlockDepth(ident, 7) = %d, want 7", got)
	}
}
