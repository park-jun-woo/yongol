//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what 정적 컨텍스트(최상위·정적 래퍼)의 data-link 요소를 파싱
package stml

import "golang.org/x/net/html"

// parseLinkStatic parses a data-link element in a static context (top
// level or inside a static wrapper) — e.g. a plain navigation link
// (`<a data-link="settings-parsing-rules">`). Children are static only;
// there is no data context to bind against.
func parseLinkStatic(n *html.Node) LinkRef {
	lr := parseLinkRef(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode || !hasContent(c) {
			continue
		}
		se := parseStaticElement(c)
		lr.Children = append(lr.Children, ChildNode{Kind: "static", Static: &se})
	}
	return lr
}
