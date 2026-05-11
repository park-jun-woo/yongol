//ff:func feature=stml-parse type=parser control=selection
//ff:what 최상위 요소를 fetch·action·static으로 분기 처리
package stml

import "golang.org/x/net/html"

// dispatchTopLevelElement handles a single top-level element node.
func dispatchTopLevelElement(n *html.Node, page *PageSpec) bool {
	switch {
	case getAttr(n, "data-fetch") != "":
		fb := parseFetchBlock(n, getAttr(n, "data-fetch"))
		page.Fetches = append(page.Fetches, fb)
		page.Children = append(page.Children, ChildNode{Kind: "fetch", Fetch: &fb})
		return true
	case getAttr(n, "data-action") != "":
		ab := parseActionBlock(n, getAttr(n, "data-action"))
		page.Actions = append(page.Actions, ab)
		page.Children = append(page.Children, ChildNode{Kind: "action", Action: &ab})
		return true
	case hasDescendantData(n):
		se := parseStaticWithDataChildren(n, page)
		page.Children = append(page.Children, ChildNode{Kind: "static", Static: &se})
		return true
	case hasContent(n):
		se := parseStaticElement(n)
		page.Children = append(page.Children, ChildNode{Kind: "static", Static: &se})
		return true
	default:
		return false
	}
}
