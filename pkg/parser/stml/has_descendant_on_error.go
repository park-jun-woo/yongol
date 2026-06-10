//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what 하위 요소에 data-on-error 속성이 있는지 재귀 확인
package stml

import "golang.org/x/net/html"

// hasDescendantOnError checks if any descendant element carries the
// data-on-error marker attribute (the error-message slot of an action block).
func hasDescendantOnError(n *html.Node) bool {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode && hasAttr(c, "data-on-error") {
			return true
		}
		if hasDescendantOnError(c) {
			return true
		}
	}
	return false
}
