//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what fetch 블록 내 정적 요소를 파싱하고 하위 data-* 요소를 부모 fb에 수집
package stml

import "golang.org/x/net/html"

// parseStaticInFetch parses a static element inside a fetch block, but still
// collects any nested data-* elements into the parent fb for validation.
func parseStaticInFetch(n *html.Node, fb *FetchBlock) StaticElement {
	se := StaticElement{
		Tag:       n.Data,
		ClassName: getAttr(n, "class"),
		Text:      directText(n),
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode || c.Type != html.ElementNode {
			continue
		}
		child := dispatchStaticFetchChild(c, fb)
		if child != nil {
			se.Children = append(se.Children, *child)
		}
	}
	return se
}
