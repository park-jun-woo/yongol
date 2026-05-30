//ff:func feature=gen-gogin type=test control=branch topic=defensive
//ff:what TestCollectResourceFromBlock — top-level + 중첩 if-블록 내 미닫힘 리소스 DF-06 재귀 수집

package qcheck

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func parseFuncBody(t *testing.T, src string) (*ast.BlockStmt, *token.FileSet) {
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

func TestCollectResourceFromBlock_NestedMissing(t *testing.T) {
	// os.Open inside an if-block with no defer Close -> one DF-06 from recursion.
	src := `package x
func H(cond bool) {
	if cond {
		f, err := os.Open("x")
		_ = err
		_ = f
	}
}`
	body, fset := parseFuncBody(t, src)
	findings := collectResourceFromBlock(body, fset)
	if len(findings) != 1 {
		t.Fatalf("want 1 DF-06 finding from nested block, got %d: %+v", len(findings), findings)
	}
	if findings[0].Category != "DF-06" || findings[0].Detail != "os.Open" {
		t.Errorf("unexpected finding: %+v", findings[0])
	}
}

func TestCollectResourceFromBlock_TopLevelClosed(t *testing.T) {
	// Resource acquired at top level WITH defer Close -> no finding.
	src := `package x
func H() {
	f, err := os.Open("x")
	defer f.Close()
	_ = err
}`
	body, fset := parseFuncBody(t, src)
	if got := collectResourceFromBlock(body, fset); len(got) != 0 {
		t.Errorf("expected no findings when closed, got %+v", got)
	}
}
