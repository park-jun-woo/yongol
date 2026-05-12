//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what walkLayoutNode — DOM 트리 재귀 순회하며 data-nav, slot data-outlet 추출

package stml

import "golang.org/x/net/html"

// walkLayoutNode recursively walks the DOM tree to extract data-nav links
// and slot data-outlet elements.
func walkLayoutNode(n *html.Node, layout *LayoutSpec) {
	if n.Type == html.ElementNode {
		collectLayoutElement(n, layout)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walkLayoutNode(c, layout)
	}
}
