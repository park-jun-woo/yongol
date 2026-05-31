//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWalkResourceBlocks — 전 함수 순회로 미닫힘 리소스 수집 + 비함수/본문없음 스킵
package qcheck

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestWalkResourceBlocks_NoFuncs(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", "package x\nvar y = 1\n", parser.SkipObjectResolution)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if got := walkResourceBlocks(file, fset); len(got) != 0 {
		t.Errorf("expected no findings, got %+v", got)
	}
}
