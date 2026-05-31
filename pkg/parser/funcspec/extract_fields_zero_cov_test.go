//ff:func feature=funcspec type=test control=sequence
//ff:what funcspec 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package funcspec

import (
	"go/ast"
	"testing"
)

func TestExtractFields_ZeroCov(t *testing.T) {
	_, f := bnFile(t, bnStructSrc)
	st := f.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Type.(*ast.StructType)
	if got := extractFields(st); len(got) != 2 {
		t.Errorf("expected 2 fields, got %d", len(got))
	}
}
