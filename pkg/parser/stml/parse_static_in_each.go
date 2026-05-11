//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what each 항목 내 정적 래퍼 요소를 파싱
package stml

import "golang.org/x/net/html"

// parseStaticInEach parses a static wrapper inside an each item.
func parseStaticInEach(n *html.Node, eb *EachBlock) StaticElement {
	se := StaticElement{
		Tag:       n.Data,
		ClassName: getAttr(n, "class"),
		Text:      directText(n),
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode {
			continue
		}
		child := dispatchStaticEachChild(c, eb)
		se.Children = append(se.Children, child)
	}
	return se
}
