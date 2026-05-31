//ff:func feature=gen-splitter type=test control=sequence
//ff:what splitter 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package splitter

import (
	"strings"
	"testing"
)

func TestCollectImportsForDecl_ZeroCov(t *testing.T) {
	_, f := bnParseFile(t, "package p\nimport (\n\t\"fmt\"\n\t\"net/http\"\n)\nfunc f() { fmt.Println(1) }")
	out := collectImportsForDecl(f.Decls[1:], f.Imports)
	if len(out) != 1 || !strings.Contains(out[0].Path.Value, "fmt") {
		t.Errorf("expected only fmt import, got %#v", out)
	}
}
