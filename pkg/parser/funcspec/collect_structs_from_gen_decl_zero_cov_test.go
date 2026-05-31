//ff:func feature=funcspec type=test control=sequence
//ff:what funcspec 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package funcspec

import (
	"go/ast"
	"testing"
)

func TestCollectStructsFromGenDecl_ZeroCov(t *testing.T) {
	_, f := bnFile(t, bnStructSrc)
	result := map[string][]Field{}
	collectStructsFromGenDecl(f.Decls[0].(*ast.GenDecl), result)
	if _, ok := result["FooRequest"]; !ok {
		t.Errorf("FooRequest not collected: %v", result)
	}
}
