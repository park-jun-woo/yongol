//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWalkUnmarshalBlocks — 전 함수 순회로 미가드 Unmarshal DF-01 수집 + 비함수/본문없음 스킵
package qcheck

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestWalkUnmarshalBlocks_NoFuncs(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", "package x\nvar y = 1\n", parser.SkipObjectResolution)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if got := walkUnmarshalBlocks(file, fset); len(got) != 0 {
		t.Errorf("expected no findings, got %+v", got)
	}
}
