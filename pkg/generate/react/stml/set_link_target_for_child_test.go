//ff:func feature=stml-gen type=test control=sequence
//ff:what TestSetLinkTargetForChild — Kind별 분기(link/fetch/each/static/state/기타) 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSetLinkTargetForChild(t *testing.T) {
	patterns := map[string]string{"p": "/p/:ID"}
	newLink := func() *stmlparser.LinkRef { return &stmlparser.LinkRef{TargetPage: "p"} }

	direct := newLink()
	setLinkTargetForChild(stmlparser.ChildNode{Kind: "link", Link: direct}, patterns)
	if direct.TargetPattern != "/p/:ID" {
		t.Errorf("link: %q", direct.TargetPattern)
	}

	inFetch := newLink()
	setLinkTargetForChild(stmlparser.ChildNode{Kind: "fetch", Fetch: &stmlparser.FetchBlock{
		Children: []stmlparser.ChildNode{{Kind: "link", Link: inFetch}},
	}}, patterns)
	if inFetch.TargetPattern != "/p/:ID" {
		t.Errorf("fetch: %q", inFetch.TargetPattern)
	}

	row := newLink()
	inEach := newLink()
	setLinkTargetForChild(stmlparser.ChildNode{Kind: "each", Each: &stmlparser.EachBlock{
		RowLink:  row,
		Children: []stmlparser.ChildNode{{Kind: "link", Link: inEach}},
	}}, patterns)
	if row.TargetPattern != "/p/:ID" || inEach.TargetPattern != "/p/:ID" {
		t.Errorf("each: row=%q child=%q", row.TargetPattern, inEach.TargetPattern)
	}

	inStatic := newLink()
	setLinkTargetForChild(stmlparser.ChildNode{Kind: "static", Static: &stmlparser.StaticElement{
		Children: []stmlparser.ChildNode{{Kind: "link", Link: inStatic}},
	}}, patterns)
	if inStatic.TargetPattern != "/p/:ID" {
		t.Errorf("static: %q", inStatic.TargetPattern)
	}

	inState := newLink()
	setLinkTargetForChild(stmlparser.ChildNode{Kind: "state", State: &stmlparser.StateBind{
		Children: []stmlparser.ChildNode{{Kind: "link", Link: inState}},
	}}, patterns)
	if inState.TargetPattern != "/p/:ID" {
		t.Errorf("state: %q", inState.TargetPattern)
	}

	// Non-container kinds are a no-op.
	setLinkTargetForChild(stmlparser.ChildNode{Kind: "bind", Bind: &stmlparser.FieldBind{Name: "x"}}, patterns)
}
