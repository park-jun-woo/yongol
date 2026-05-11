//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what action 블록 내부를 순회하며 필드·submit·component 수집
package stml

import "golang.org/x/net/html"

// walkActionChildren recursively collects fields inside an action block.
func walkActionChildren(n *html.Node, ab *ActionBlock) {
	if n.Type != html.ElementNode {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walkActionChildren(c, ab)
		}
		return
	}

	dispatched := dispatchActionChild(n, ab)
	if dispatched {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walkActionChildren(c, ab)
	}
}
