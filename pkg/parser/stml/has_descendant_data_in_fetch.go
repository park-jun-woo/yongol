//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what 하위 요소에 data-* 속성이 있는지 재귀 확인 (fetch 내부용)
package stml

import "golang.org/x/net/html"

// hasDescendantDataInFetch checks if any descendant has data-* attributes relevant to fetch.
func hasDescendantDataInFetch(n *html.Node) bool {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && (hasDataAttr(c) || hasDescendantDataInFetch(c)) {
			return true
		}
	}
	return false
}
