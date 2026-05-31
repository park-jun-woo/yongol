//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestCollectFileLoopReports — 여러 FuncDecl의 루프 집계 + 비함수/본문없음 스킵 검증
package qcheck

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestCollectFileLoopReports(t *testing.T) {
	src := `package x

var g = 1

func decl()

func a() {
	for i := 0; i < 2; i++ {
		_ = i
	}
}

func b() {
	for _, v := range []int{1} {
		_ = v
	}
	for j := 0; j < 1; j++ {
		_ = j
	}
}
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	reports := collectFileLoopReports(fset, file, src)
	// a:1 loop, b:2 loops -> 3 total.
	if len(reports) != 3 {
		t.Fatalf("want 3 loop reports, got %d: %+v", len(reports), reports)
	}
	byFunc := map[string]int{}
	for _, r := range reports {
		byFunc[r.Func]++
	}
	if byFunc["a"] != 1 || byFunc["b"] != 2 {
		t.Errorf("per-func loop counts = %v, want a:1 b:2", byFunc)
	}
}
