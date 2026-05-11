//ff:func feature=stml-parse type=parser control=selection
//ff:what 정적 요소의 data-* 자식을 분기 처리
package stml

import "golang.org/x/net/html"

func dispatchStaticDataChild(c *html.Node, page *PageSpec) *ChildNode {
	switch {
	case getAttr(c, "data-fetch") != "":
		fb := parseFetchBlock(c, getAttr(c, "data-fetch"))
		page.Fetches = append(page.Fetches, fb)
		return &ChildNode{Kind: "fetch", Fetch: &fb}
	case getAttr(c, "data-action") != "":
		ab := parseActionBlock(c, getAttr(c, "data-action"))
		page.Actions = append(page.Actions, ab)
		return &ChildNode{Kind: "action", Action: &ab}
	case hasDescendantData(c):
		child := parseStaticWithDataChildren(c, page)
		return &ChildNode{Kind: "static", Static: &child}
	case hasContent(c):
		child := parseStaticElement(c)
		return &ChildNode{Kind: "static", Static: &child}
	default:
		return nil
	}
}
