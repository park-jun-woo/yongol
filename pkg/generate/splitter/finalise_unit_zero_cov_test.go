//ff:func feature=gen-splitter type=test control=sequence
//ff:what splitter 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package splitter

import (
	"testing"
)

func TestFinaliseUnit_ZeroCov(t *testing.T) {
	_, f := bnParseFile(t, "package p\nimport \"fmt\"\nfunc Foo() { fmt.Println(1) }")
	u := &splitUnit{
		PkgName: "p",
		Decls:   f.Decls[1:],
		Docs:    []string{"Foo does X."},
	}
	finaliseUnit(u, "auth", "handler", nil, f.Imports)
	if len(u.Annotations) == 0 {
		t.Errorf("annotations not filled")
	}
}
