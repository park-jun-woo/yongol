//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what 요소에 텍스트 또는 자식 요소가 있는지 확인
package stml

import (
	"strings"

	"golang.org/x/net/html"
)

// hasContent returns true if the element has text or element children.
func hasContent(n *html.Node) bool {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode && strings.TrimSpace(c.Data) != "" {
			return true
		}
		if c.Type == html.ElementNode {
			return true
		}
	}
	return false
}
