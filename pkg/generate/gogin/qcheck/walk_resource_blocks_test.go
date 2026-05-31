//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWalkResourceBlocks — 전 함수 순회로 미닫힘 리소스 수집 + 비함수/본문없음 스킵
package qcheck

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestWalkResourceBlocks(t *testing.T) {
	src := `package x

var g = 1

func decl()

func ok() {
	f, err := os.Open(p)
	defer f.Close()
	_ = err
}

func leak() {
	f, err := os.Open(p)
	_ = err
	_ = f
}`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", src, parser.SkipObjectResolution)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	findings := walkResourceBlocks(file, fset)
	if len(findings) != 1 {
		t.Fatalf("want 1 finding (only leak()), got %d: %+v", len(findings), findings)
	}
	if findings[0].Category != "DF-06" {
		t.Errorf("unexpected finding: %+v", findings[0])
	}
}
