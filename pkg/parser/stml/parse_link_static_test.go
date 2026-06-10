//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParseLinkStatic — 정적 컨텍스트(최상위·정적 래퍼)의 data-link 파싱 검증

package stml

import (
	"strings"
	"testing"
)

func TestParseLinkStatic(t *testing.T) {
	input := `<main>
  <nav>
    <a data-link="settings-parsing-rules">파싱 규칙</a>
  </nav>
</main>`

	page, diags := ParseReader("settings.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}

	// main → static wrapper → nav (static with data-link descendant) → link.
	var link *LinkRef
	var walk func(children []ChildNode)
	walk = func(children []ChildNode) {
		for _, ch := range children {
			switch ch.Kind {
			case "link":
				link = ch.Link
			case "static":
				walk(ch.Static.Children)
			}
		}
	}
	walk(page.Children)
	if link == nil {
		t.Fatalf("no link found in page children: %+v", page.Children)
	}
	if link.TargetPage != "settings-parsing-rules" {
		t.Errorf("TargetPage = %q", link.TargetPage)
	}
	if link.Text != "파싱 규칙" {
		t.Errorf("Text = %q", link.Text)
	}
	if len(link.Params) != 0 {
		t.Errorf("Params = %+v, want none", link.Params)
	}

	// A static child inside the link element is preserved.
	el := firstElementNode(t, `<a data-link="p"><strong>go</strong></a>`, "a")
	lr := parseLinkStatic(el)
	if len(lr.Children) != 1 || lr.Children[0].Kind != "static" || lr.Children[0].Static.Tag != "strong" {
		t.Errorf("static link children = %+v", lr.Children)
	}
}
