//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what 요소에 지정한 속성이 존재하는지 확인
package stml

import "golang.org/x/net/html"

// hasAttr returns true if the element has the named attribute (regardless of value).
func hasAttr(n *html.Node, key string) bool {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return true
		}
	}
	return false
}
