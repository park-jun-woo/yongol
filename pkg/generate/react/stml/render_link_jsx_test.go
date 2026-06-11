//ff:func feature=stml-gen type=test control=sequence
//ff:what TestRenderLinkJSX — 텍스트 단일행·자식 포함 Link JSX 렌더 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderLinkJSX(t *testing.T) {
	// Text-only static link renders on a single line.
	got := renderLinkJSX(stmlparser.LinkRef{
		TargetPage:    "settings-parsing-rules",
		TargetPattern: "/settings-parsing-rules",
		Text:          "파싱 규칙",
	}, "", "item", 0, nil, bindCtx{})
	want := `<Link to="/settings-parsing-rules">파싱 규칙</Link>`
	if got != want {
		t.Errorf("text link:\n got %q\nwant %q", got, want)
	}

	// A link with a bind child wraps the rendered children; direct text
	// is preserved before them.
	bind := stmlparser.FieldBind{Name: "name", Tag: "span"}
	got = renderLinkJSX(stmlparser.LinkRef{
		TargetPage:    "building-detail",
		TargetPattern: "/buildings/:BuildingID",
		Params:        []stmlparser.LinkParamBind{{Source: "item.id", Segment: "BuildingID"}},
		Text:          "보기",
		Children:      []stmlparser.ChildNode{{Kind: "bind", Bind: &bind}},
	}, "data", "item", 0, nil, bindCtx{})
	assertContains(t, got, "<Link to={`/buildings/${item.id}`}>")
	assertContains(t, got, "보기")
	assertContains(t, got, "</Link>")
}
