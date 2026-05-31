//ff:func feature=funcspec type=test control=iteration dimension=1
//ff:what funcspec 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package funcspec

import (
	"go/ast"
	"testing"
)

func TestProcessTypeSpecs_ZeroCov(t *testing.T) {
	_, f := bnFile(t, bnStructSrc)
	spec := &FuncSpec{Name: "foo"}
	for _, d := range f.Decls {
		if gd, ok := d.(*ast.GenDecl); ok {
			processTypeSpecs(gd, spec, "FooRequest", "FooResponse")
		}
	}
	if len(spec.RequestFields) != 2 || len(spec.ResponseFields) != 1 {
		t.Errorf("type specs not processed: %#v", spec)
	}
}
