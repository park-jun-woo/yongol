//ff:func feature=ssac-parse type=test-helper control=iteration dimension=1
//ff:what parseSSaCSource — SSaC 테스트용 함수 선행 주석 수집 헬퍼

package ssac

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

// parseSSaCSource parses src and returns the leading comments attached
// to the first FuncDecl (mirrors collectFuncComments so tests exercise
// the production comment surface).
func parseSSaCSource(t *testing.T, src string) []*ast.Comment {
	t.Helper()
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "x.ssac", src, parser.ParseComments)
	if err != nil {
		t.Fatalf("parse: %v", err)
	}
	for _, d := range f.Decls {
		fn, ok := d.(*ast.FuncDecl)
		if !ok {
			continue
		}
		return collectFuncComments(f, fn.Pos())
	}
	t.Fatalf("no FuncDecl in source")
	return nil
}
