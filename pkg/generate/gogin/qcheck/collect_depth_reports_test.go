//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestCollectDepthReports — FuncDecl만 DepthReport로, var/외부선언/본문없음 스킵 검증
package qcheck

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestCollectDepthReports(t *testing.T) {
	src := `package x

var g = 1

func noBody()

func a() {
	if true {
		_ = 1
	}
}

func b() { _ = 2 }
`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	reports := collectDepthReports(file)
	// noBody (no body) and var skipped -> only a and b.
	if len(reports) != 2 {
		t.Fatalf("want 2 reports, got %d: %+v", len(reports), reports)
	}
	names := map[string]bool{}
	for _, r := range reports {
		names[r.Func] = true
	}
	if !names["a"] || !names["b"] {
		t.Errorf("expected funcs a and b, got %v", names)
	}
}
