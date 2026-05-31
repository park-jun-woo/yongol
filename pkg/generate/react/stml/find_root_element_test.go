//ff:func feature=stml-gen type=test control=sequence
//ff:what render_root_params_unit_test — findRootElement / renderUseParams 단위 테스트
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestFindRootElement(t *testing.T) {
	// single static child -> use its tag/class
	single := stmlparser.PageSpec{
		Children: []stmlparser.ChildNode{
			{Kind: "static", Static: &stmlparser.StaticElement{Tag: "main", ClassName: "container"}},
		},
	}
	if tag, cls := findRootElement(single); tag != "main" || cls != "container" {
		t.Errorf("single static -> (%q,%q), want (main, container)", tag, cls)
	}

	// multiple children -> default div with empty class
	multi := stmlparser.PageSpec{
		Children: []stmlparser.ChildNode{
			{Kind: "static", Static: &stmlparser.StaticElement{Tag: "header"}},
			{Kind: "static", Static: &stmlparser.StaticElement{Tag: "footer"}},
		},
	}
	if tag, cls := findRootElement(multi); tag != "div" || cls != "" {
		t.Errorf("multi child -> (%q,%q), want (div, \"\")", tag, cls)
	}

	// single non-static child -> default div
	nonStatic := stmlparser.PageSpec{
		Children: []stmlparser.ChildNode{
			{Kind: "fetch", Fetch: &stmlparser.FetchBlock{}},
		},
	}
	if tag, cls := findRootElement(nonStatic); tag != "div" || cls != "" {
		t.Errorf("single non-static -> (%q,%q), want (div, \"\")", tag, cls)
	}

	// no children -> default div
	if tag, cls := findRootElement(stmlparser.PageSpec{}); tag != "div" || cls != "" {
		t.Errorf("empty -> (%q,%q), want (div, \"\")", tag, cls)
	}
}
