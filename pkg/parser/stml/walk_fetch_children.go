//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what fetch 블록 내부를 순회하며 바인딩·each·state·component 수집
package stml

import "golang.org/x/net/html"

// walkFetchChildren recursively collects bindings inside a fetch block.
func walkFetchChildren(n *html.Node, fb *FetchBlock) {
	if n.Type == html.TextNode {
		return
	}
	if n.Type != html.ElementNode {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			walkFetchChildren(c, fb)
		}
		return
	}

	dispatched := dispatchFetchChild(n, fb)
	if dispatched {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walkFetchChildren(c, fb)
	}
}
