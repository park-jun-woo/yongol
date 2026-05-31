//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestWalkScanBlocks — 전 함수 순회로 미가드 .Scan() DF-02 수집 + 비함수/본문없음 스킵
package qcheck

import (
	"go/parser"
	"go/token"
	"testing"
)

func TestWalkScanBlocks(t *testing.T) {
	src := `package x

func decl()

func ok(r row) {
	err := r.Scan(nil)
	if err != nil { return }
}

func bad(r row) {
	_ = r.Scan(nil)
}`
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", src, parser.SkipObjectResolution)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	findings := walkScanBlocks(file, fset)
	if len(findings) == 0 {
		t.Fatalf("want at least 1 DF-02 finding from bad(), got none")
	}
	for _, f := range findings {
		if f.Category != "DF-02" {
			t.Errorf("unexpected finding category: %+v", f)
		}
	}
}
