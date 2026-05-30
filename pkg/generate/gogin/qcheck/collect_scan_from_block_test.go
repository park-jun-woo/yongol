//ff:func feature=gen-gogin type=test control=branch topic=defensive
//ff:what TestCollectScanFromBlock — 중첩 블록 내 미가드 .Scan() DF-02 재귀 수집 + 가드 케이스

package qcheck

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func parseHBody(t *testing.T, src string) (*ast.BlockStmt, *token.FileSet) {
	t.Helper()
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "f.go", src, 0)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	for _, d := range file.Decls {
		if fn, ok := d.(*ast.FuncDecl); ok && fn.Name.Name == "H" {
			return fn.Body, fset
		}
	}
	t.Fatalf("func H not found")
	return nil, nil
}

func TestCollectScanFromBlock_NestedUnguarded(t *testing.T) {
	src := `package x
func H(cond bool, r row) {
	if cond {
		_ = r.Scan(nil)
	}
}`
	body, fset := parseHBody(t, src)
	findings := collectScanFromBlock(body, fset)
	if len(findings) == 0 {
		t.Fatalf("want at least 1 DF-02 finding from nested block, got none")
	}
	for _, f := range findings {
		if f.Category != "DF-02" {
			t.Errorf("unexpected finding category: %+v", f)
		}
	}
}

func TestCollectScanFromBlock_Guarded(t *testing.T) {
	src := `package x
func H(r row) {
	err := r.Scan(nil)
	if err != nil { return }
}`
	body, fset := parseHBody(t, src)
	if got := collectScanFromBlock(body, fset); len(got) != 0 {
		t.Errorf("expected no findings for guarded Scan, got %+v", got)
	}
}
