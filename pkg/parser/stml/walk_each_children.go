//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what each 블록 내부를 순회하며 바인딩·state·component 수집
package stml

import "golang.org/x/net/html"

// walkEachChildren recursively collects bindings inside an each block's item.
func walkEachChildren(n *html.Node, eb *EachBlock) {
	if n.Type != html.ElementNode {
		return
	}

	dispatched := dispatchEachChild(n, eb)
	if dispatched {
		return
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		walkEachChildren(c, eb)
	}
}
