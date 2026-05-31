//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestCollectDepthReports — FuncDecl만 DepthReport로, var/외부선언/본문없음 스킵 검증
package qcheck

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestCollectDepthReports_Empty(t *testing.T) {
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", "package x\nvar y = 1\n", 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	if got := collectDepthReports(file); len(got) != 0 {
		t.Errorf("expected no reports, got %+v", got)
	}
}
