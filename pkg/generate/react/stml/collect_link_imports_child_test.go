//ff:func feature=stml-gen type=test control=iteration dimension=1
//ff:what TestCollectLinkImportsChild — Kind별 분기(link/fetch/each/static/state/기타) 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectLinkImportsChild(t *testing.T) {
	link := stmlparser.LinkRef{TargetPage: "p"}
	rowLink := stmlparser.LinkRef{TargetPage: "q", Params: []stmlparser.LinkParamBind{{Source: "route.ID"}}}
	linkNode := stmlparser.ChildNode{Kind: "link", Link: &link}
	cases := []stmlparser.ChildNode{
		linkNode,
		{Kind: "fetch", Fetch: &stmlparser.FetchBlock{Children: []stmlparser.ChildNode{linkNode}}},
		{Kind: "each", Each: &stmlparser.EachBlock{RowLink: &rowLink}},
		{Kind: "static", Static: &stmlparser.StaticElement{Children: []stmlparser.ChildNode{linkNode}}},
		{Kind: "state", State: &stmlparser.StateBind{Children: []stmlparser.ChildNode{linkNode}}},
		{Kind: "bind", Bind: &stmlparser.FieldBind{Name: "x"}},
	}
	for _, ch := range cases {
		var is importSet
		collectLinkImportsChild(ch, &is)
		wantLink := ch.Kind != "bind"
		if is.useLink != wantLink {
			t.Errorf("kind %s: useLink=%v want %v", ch.Kind, is.useLink, wantLink)
		}
		if ch.Kind == "each" && !is.useParams {
			t.Error("each RowLink with route.* source must set useParams")
		}
	}
}
