//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what 하위 요소에 data-field 속성이 있는지 재귀 확인
package stml

import "golang.org/x/net/html"

// hasDescendantField checks if any descendant has data-field or data-component with data-field.
func hasDescendantField(n *html.Node) bool {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && (hasFieldAttr(c) || hasDescendantField(c)) {
			return true
		}
	}
	return false
}
