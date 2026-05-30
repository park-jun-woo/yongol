//ff:func feature=gen-gogin type=test control=sequence topic=loop-report
//ff:what TestCollectLoopReports — 한 FuncDecl 내 중첩 포함 모든 루프를 리포트로 변환 검증

package qcheck

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestCollectLoopReports(t *testing.T) {
	src := `package x
func work() {
	for i := 0; i < 2; i++ {
		for _, v := range []int{1} {
			_ = v
		}
	}
}`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	fn := file.Decls[0].(*ast.FuncDecl)
	reports := collectLoopReports(fset, fn, src)
	// outer for + inner range = 2.
	if len(reports) != 2 {
		t.Fatalf("want 2 reports, got %d: %+v", len(reports), reports)
	}
	for _, r := range reports {
		if r.Func != "work" {
			t.Errorf("Func = %q, want work", r.Func)
		}
	}
}

func TestCollectLoopReports_NoLoop(t *testing.T) {
	src := "package x\nfunc f() { _ = 1 }"
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	fn := file.Decls[0].(*ast.FuncDecl)
	if got := collectLoopReports(fset, fn, src); len(got) != 0 {
		t.Errorf("expected no reports, got %+v", got)
	}
}
