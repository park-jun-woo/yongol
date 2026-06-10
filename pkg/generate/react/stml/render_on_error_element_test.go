//ff:func feature=stml-gen type=test control=sequence
//ff:what renderOnErrorElement — 에러 상태 조건부 렌더 JSX(className 유무) 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderOnErrorElement(t *testing.T) {
	// with className -> className attribute emitted, indented
	withCls := stmlparser.StaticElement{Tag: "p", ClassName: "error"}
	if got, want := renderOnErrorElement(withCls, "loginError", 6),
		`      {loginError && <p className="error">{loginError}</p>}`; got != want {
		t.Errorf("with className = %q, want %q", got, want)
	}

	// no className -> no attribute, no indent
	noCls := stmlparser.StaticElement{Tag: "span"}
	if got, want := renderOnErrorElement(noCls, "loginError", 0),
		`{loginError && <span>{loginError}</span>}`; got != want {
		t.Errorf("no className = %q, want %q", got, want)
	}
}
