//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestWalkScanBlocks — 전 함수 순회로 미가드 .Scan() DF-02 수집 + 비함수/본문없음 스킵
package qcheck

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestWalkScanBlocks_NoFuncs(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", "package x\nvar y = 1\n", parser.SkipObjectResolution)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if got := walkScanBlocks(file, fset); len(got) != 0 {
		t.Errorf("expected no findings, got %+v", got)
	}
}
