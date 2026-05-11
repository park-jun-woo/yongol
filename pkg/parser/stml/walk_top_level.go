//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what DOM 트리를 순회하며 최상위 블록(fetch·action·static) 수집
package stml

import "golang.org/x/net/html"

// walkTopLevel traverses the DOM tree collecting top-level blocks.
func walkTopLevel(n *html.Node, page *PageSpec) {
	if n.Type == html.ElementNode && isImplicitTag(n.Data) {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walkTopLevel(c, page)
		}
		return
	}
	if n.Type == html.ElementNode {
		if dispatchTopLevelElement(n, page) {
			return
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walkTopLevel(c, page)
	}
}
