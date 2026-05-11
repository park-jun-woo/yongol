//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what 바인딩이 없는 정적 HTML 요소를 재귀적으로 파싱
package stml

import "golang.org/x/net/html"

// parseStaticElement recursively parses a non-binding HTML element.
func parseStaticElement(n *html.Node) StaticElement {
	se := StaticElement{
		Tag:       n.Data,
		ClassName: getAttr(n, "class"),
		Text:      directText(n),
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && hasContent(c) {
			child := parseStaticElement(c)
			se.Children = append(se.Children, ChildNode{Kind: "static", Static: &child})
		}
	}
	return se
}
