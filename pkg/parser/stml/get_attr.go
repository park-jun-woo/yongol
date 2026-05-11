//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what HTML 요소에서 지정한 속성 값을 반환
package stml

import "golang.org/x/net/html"

// getAttr returns the value of the named attribute, or "" if not found.
func getAttr(n *html.Node, key string) string {
	for _, attr := range n.Attr {
		if attr.Key == key {
			return attr.Val
		}
	}
	return ""
}
