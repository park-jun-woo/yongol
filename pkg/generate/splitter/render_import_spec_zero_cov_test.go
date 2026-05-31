//ff:func feature=gen-splitter type=test control=sequence
//ff:what splitter 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package splitter

import (
	"go/ast"
	"testing"
)

func TestRenderImportSpec_ZeroCov(t *testing.T) {
	plain := &ast.ImportSpec{Path: &ast.BasicLit{Value: `"fmt"`}}
	if got := renderImportSpec(plain); got != `"fmt"` {
		t.Errorf("plain=%q", got)
	}
	aliased := &ast.ImportSpec{Name: &ast.Ident{Name: "f"}, Path: &ast.BasicLit{Value: `"fmt"`}}
	if got := renderImportSpec(aliased); got != `f "fmt"` {
		t.Errorf("aliased=%q", got)
	}
}
