//ff:func feature=stml-gen type=test control=sequence
//ff:what render_component_static_unit_test — renderComponentJSX / renderStaticActionJSX 단위 테스트
package stml

import (
	"strings"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderStaticActionJSX(t *testing.T) {
	// self-closing (no children, no text)
	got := renderStaticActionJSX(stmlparser.StaticElement{Tag: "hr"}, "form", 2)
	if got != "  <hr />" {
		t.Errorf("self-closing = %q, want '  <hr />'", got)
	}

	// text content with class
	got = renderStaticActionJSX(stmlparser.StaticElement{Tag: "label", ClassName: "lbl", Text: "Email"}, "form", 0)
	if got != `<label className="lbl">Email</label>` {
		t.Errorf("text+class = %q", got)
	}

	// nested children -> multi-line open/close with indented children
	parent := stmlparser.StaticElement{
		Tag:      "div",
		Children: []stmlparser.ChildNode{{Kind: "static", Static: &stmlparser.StaticElement{Tag: "span", Text: "hi"}}},
	}
	got = renderStaticActionJSX(parent, "form", 2)
	lines := strings.Split(got, "\n")
	if len(lines) != 3 {
		t.Fatalf("expected 3 lines, got %d: %q", len(lines), got)
	}
	if lines[0] != "  <div>" || lines[2] != "  </div>" {
		t.Errorf("open/close lines wrong: %q", got)
	}
	if !strings.Contains(lines[1], "<span>hi</span>") {
		t.Errorf("child not rendered/indented: %q", lines[1])
	}
}
