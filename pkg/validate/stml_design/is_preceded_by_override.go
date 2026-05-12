//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-design
//ff:what isPrecededByOverride — 요소 앞에 @override 주석이 있는지 확인
package stml_design

import (
	"strings"

	"golang.org/x/net/html"
)

// isPrecededByOverride checks if the given element node has a preceding
// sibling that is an @override comment (skipping whitespace text nodes).
func isPrecededByOverride(n *html.Node) bool {
	for prev := n.PrevSibling; prev != nil; prev = prev.PrevSibling {
		if prev.Type == html.CommentNode && isOverrideComment(prev.Data) {
			return true
		}
		// Skip whitespace text nodes
		if prev.Type == html.TextNode && strings.TrimSpace(prev.Data) == "" {
			continue
		}
		break
	}
	return false
}
