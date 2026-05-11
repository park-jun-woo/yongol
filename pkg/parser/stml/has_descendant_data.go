//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what 하위 요소에 data-fetch 또는 data-action 속성이 있는지 재귀 확인
package stml

import "golang.org/x/net/html"

// hasDescendantData checks if any descendant has a data-fetch or data-action attribute.
func hasDescendantData(n *html.Node) bool {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if isDataElement(c) || (c.Type == html.ElementNode && hasDescendantData(c)) {
			return true
		}
	}
	return false
}
