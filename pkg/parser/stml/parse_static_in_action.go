//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what action 블록 내 정적 요소를 파싱하고 하위 필드·컴포넌트 수집
package stml

import "golang.org/x/net/html"

// parseStaticInAction parses a static element inside an action block.
func parseStaticInAction(n *html.Node, ab *ActionBlock) StaticElement {
	se := StaticElement{
		Tag:       n.Data,
		ClassName: getAttr(n, "class"),
		Text:      directText(n),
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type != html.ElementNode {
			continue
		}
		child := dispatchStaticActionChild(c, ab)
		if child != nil {
			se.Children = append(se.Children, *child)
		}
	}
	return se
}
