//ff:func feature=gen-splitter type=test control=sequence
//ff:what splitter 내부 함수 by-name 직접 호출 테스트 — tsma content-aware 귀속용
package splitter

import (
	"strings"
	"testing"
)

func TestInjectFFAnnotation_ZeroCov(t *testing.T) {
	fn := firstDecl(t, "package p\n// Foo does X.\nfunc Foo() {}")
	lines := injectFFAnnotation(fn, "auth", "handler", "sequence", "", "Foo does X.")
	if len(lines) != 2 || !strings.HasPrefix(lines[0], "ff:func") {
		t.Errorf("func annotation wrong: %#v", lines)
	}
	td := firstDecl(t, "package p\ntype Bar struct{}")
	tlines := injectFFAnnotation(td, "auth", "model", "sequence", "", "")
	if !strings.HasPrefix(tlines[0], "ff:type") {
		t.Errorf("type annotation wrong: %#v", tlines)
	}
}
