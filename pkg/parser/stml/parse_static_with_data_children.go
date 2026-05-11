//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what data-* 하위 요소를 가진 정적 요소를 파싱
package stml

import "golang.org/x/net/html"

// parseStaticWithDataChildren parses a static element that has data-* descendants.
func parseStaticWithDataChildren(n *html.Node, page *PageSpec) StaticElement {
	se := StaticElement{
		Tag:       n.Data,
		ClassName: getAttr(n, "class"),
		Text:      directText(n),
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode {
			continue
		}
		child := dispatchStaticDataChild(c, page)
		if child != nil {
			se.Children = append(se.Children, *child)
		}
	}
	return se
}
