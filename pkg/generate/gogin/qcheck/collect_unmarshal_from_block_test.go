//ff:func feature=gen-gogin type=test control=branch topic=defensive
//ff:what TestCollectUnmarshalFromBlock — 중첩 블록 내 미가드 json.Unmarshal DF-01 재귀 + 가드 케이스

package qcheck

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func unmarshalBody(t *testing.T, src string) (*ast.BlockStmt, *token.FileSet) {
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

func TestCollectUnmarshalFromBlock_NestedUnguarded(t *testing.T) {
	src := `package x
func H(cond bool, b []byte, v any) {
	if cond {
		_ = json.Unmarshal(b, v)
	}
}`
	body, fset := unmarshalBody(t, src)
	findings := collectUnmarshalFromBlock(body, []string{"json"}, fset)
	if len(findings) == 0 {
		t.Fatalf("want at least 1 DF-01 finding from nested block, got none")
	}
	for _, f := range findings {
		if f.Category != "DF-01" || f.Detail != "json.Unmarshal" {
			t.Errorf("unexpected finding: %+v", f)
		}
	}
}

func TestCollectUnmarshalFromBlock_Guarded(t *testing.T) {
	src := `package x
func H(b []byte, v any) {
	if err := json.Unmarshal(b, v); err != nil { return }
}`
	body, fset := unmarshalBody(t, src)
	if got := collectUnmarshalFromBlock(body, []string{"json"}, fset); len(got) != 0 {
		t.Errorf("expected no findings for guarded Unmarshal, got %+v", got)
	}
}
