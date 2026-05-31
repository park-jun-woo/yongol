//ff:func feature=funcspec type=test control=sequence
//ff:what funcspec 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package funcspec

import (
	"go/ast"
	"testing"
)

func TestBuildFields_ZeroCov(t *testing.T) {
	_, f := bnFile(t, bnStructSrc)
	gd := f.Decls[0].(*ast.GenDecl)
	st := gd.Specs[0].(*ast.TypeSpec).Type.(*ast.StructType)
	fields := buildFields(st.Fields.List[0])
	if len(fields) != 1 || fields[0].Name != "Name" || fields[0].JSONName != "name" {
		t.Errorf("buildFields wrong: %#v", fields)
	}
}
