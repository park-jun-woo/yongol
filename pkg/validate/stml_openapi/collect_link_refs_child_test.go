//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what collectLinkRefsChild — Kind별 분기(link/fetch/each/static/state)와 컨텍스트 리셋 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectLinkRefsChild(t *testing.T) {
	raif := map[string]map[string]map[string]bool{
		"ListB": {"items": {"id": true}},
	}
	link := stml.LinkRef{TargetPage: "p"}
	linkNode := stml.ChildNode{Kind: "link", Link: &link}

	// link inherits the passed context.
	var refs []linkRefCtx
	collectLinkRefsChild(linkNode, "ListB", map[string]bool{"id": true}, true, raif, &refs)
	if len(refs) != 1 || !refs[0].InEach || refs[0].ItemFields == nil {
		t.Fatalf("link: %+v", refs)
	}

	// fetch resets the context (no row in scope).
	refs = nil
	collectLinkRefsChild(stml.ChildNode{Kind: "fetch", Fetch: &stml.FetchBlock{
		OperationID: "ListB",
		Children:    []stml.ChildNode{linkNode},
	}}, "", map[string]bool{"id": true}, true, raif, &refs)
	if len(refs) != 1 || refs[0].InEach || refs[0].ItemFields != nil {
		t.Fatalf("fetch reset: %+v", refs)
	}

	// each resolves the item schema and records the RowLink.
	row := stml.LinkRef{TargetPage: "q"}
	refs = nil
	collectLinkRefsChild(stml.ChildNode{Kind: "each", Each: &stml.EachBlock{
		Field:    "items",
		RowLink:  &row,
		Children: []stml.ChildNode{linkNode},
	}}, "ListB", nil, false, raif, &refs)
	if len(refs) != 2 {
		t.Fatalf("each: %+v", refs)
	}
	for _, r := range refs {
		if !r.InEach || r.ItemFields == nil || !r.ItemFields["id"] {
			t.Errorf("each ref ctx: %+v", r)
		}
	}

	// static and state wrappers pass the context through.
	refs = nil
	collectLinkRefsChild(stml.ChildNode{Kind: "static", Static: &stml.StaticElement{
		Children: []stml.ChildNode{linkNode},
	}}, "", nil, false, raif, &refs)
	collectLinkRefsChild(stml.ChildNode{Kind: "state", State: &stml.StateBind{
		Children: []stml.ChildNode{linkNode},
	}}, "", nil, false, raif, &refs)
	if len(refs) != 2 {
		t.Fatalf("static/state: %+v", refs)
	}

	// non-container kinds are a no-op.
	refs = nil
	collectLinkRefsChild(stml.ChildNode{Kind: "bind", Bind: &stml.FieldBind{Name: "x"}}, "", nil, false, raif, &refs)
	if len(refs) != 0 {
		t.Fatalf("bind: %+v", refs)
	}
}
