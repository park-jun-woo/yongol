//ff:func feature=stml-gen type=test control=sequence
//ff:what TestCollectAllLinks — 행 자식 트리의 링크 수집(직접·static 중첩, action 비재귀) 검증
package stml

import (
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectAllLinks(t *testing.T) {
	direct := stmlparser.LinkRef{TargetPage: "a"}
	nested := stmlparser.LinkRef{TargetPage: "b"}
	inState := stmlparser.LinkRef{TargetPage: "c"}
	wrapper := stmlparser.StaticElement{
		Tag:      "div",
		Children: []stmlparser.ChildNode{{Kind: "link", Link: &nested}},
	}
	stateWrapper := stmlparser.StateBind{
		Children: []stmlparser.ChildNode{{Kind: "link", Link: &inState}},
	}
	bind := stmlparser.FieldBind{Name: "name"}
	nodes := []stmlparser.ChildNode{
		{Kind: "bind", Bind: &bind},
		{Kind: "link", Link: &direct},
		{Kind: "static", Static: &wrapper},
		{Kind: "state", State: &stateWrapper},
	}
	links := collectAllLinks(nodes)
	if len(links) != 3 || links[0].TargetPage != "a" || links[1].TargetPage != "b" || links[2].TargetPage != "c" {
		t.Errorf("got %+v", links)
	}
	if got := collectAllLinks(nil); len(got) != 0 {
		t.Errorf("nil nodes: got %+v", got)
	}
}
