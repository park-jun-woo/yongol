//ff:func feature=gen-gogin type=test control=branch topic=loop-report
//ff:what TestAppendLoopReport — for/range 노드는 append, 비루프 노드는 무시 검증

package qcheck

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestAppendLoopReport(t *testing.T) {
	src := `package x
func f() {
	for i := 0; i < 3; i++ {
		_ = i
	}
	for _, v := range []int{1, 2} {
		_ = v
	}
	if true {
		_ = 1
	}
}`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}

	var out []PureLinesReport
	ast.Inspect(file, func(n ast.Node) bool {
		if n != nil {
			appendLoopReport(&out, fset, "f", n, src)
		}
		return true
	})

	if len(out) != 2 {
		t.Fatalf("want 2 loop reports, got %d", len(out))
	}
	kinds := map[string]bool{}
	for _, r := range out {
		kinds[r.LoopKind] = true
		if r.Func != "f" {
			t.Errorf("Func = %q, want f", r.Func)
		}
	}
	if !kinds["for"] || !kinds["range"] {
		t.Errorf("expected both for and range kinds, got %v", kinds)
	}
}
