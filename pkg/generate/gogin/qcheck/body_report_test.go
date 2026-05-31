//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestBodyReport — 루프 body 순수 라인 수와 헤더 라인을 PureLinesReport로 반환 검증
package qcheck

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestBodyReport(t *testing.T) {
	src := `package x
func f() {
	for i := 0; i < 3; i++ {
		a := 1
		// comment line — skipped

		b := a + 1
		_ = b
	}
}`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	var loop *ast.ForStmt
	ast.Inspect(file, func(n ast.Node) bool {
		if fs, ok := n.(*ast.ForStmt); ok {
			loop = fs
			return false
		}
		return true
	})
	if loop == nil {
		t.Fatalf("no for loop found")
	}

	rep := bodyReport(fset, "f", "for", loop.For, loop.Body, src)
	if rep.Func != "f" {
		t.Errorf("Func = %q, want f", rep.Func)
	}
	if rep.LoopKind != "for" {
		t.Errorf("LoopKind = %q, want for", rep.LoopKind)
	}
	if rep.Line != 3 {
		t.Errorf("Line = %d, want 3 (loop header)", rep.Line)
	}
	// Pure lines: a:=1, b:=a+1, _=b -> 3 (blank + comment skipped).
	if rep.PureLines != 3 {
		t.Errorf("PureLines = %d, want 3", rep.PureLines)
	}
}
