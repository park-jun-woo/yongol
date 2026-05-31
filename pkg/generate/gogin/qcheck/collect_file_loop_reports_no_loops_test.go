//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCollectFileLoopReports — 여러 FuncDecl의 루프 집계 + 비함수/본문없음 스킵 검증
package qcheck

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestCollectFileLoopReports_NoLoops(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", "package x\nfunc a() { _ = 1 }\n", 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if got := collectFileLoopReports(fset, file, "package x\nfunc a() { _ = 1 }\n"); len(got) != 0 {
		t.Errorf("expected no reports, got %+v", got)
	}
}
