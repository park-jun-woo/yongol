//ff:func feature=funcspec type=test control=iteration dimension=1
//ff:what funcspec 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package funcspec

import (
	"testing"
)

func TestProcessDecl_ZeroCov(t *testing.T) {
	fset, f := bnFile(t, bnStructSrc)
	spec := &FuncSpec{Name: "foo"}
	for _, d := range f.Decls {
		processDecl(d, fset, spec, "FooRequest", "FooResponse")
	}
	if len(spec.RequestFields) != 2 {
		t.Errorf("processDecl did not handle type decl: %#v", spec)
	}
}
