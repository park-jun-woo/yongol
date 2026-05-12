//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-design
//ff:what walkForOverrides — DOM 순회하며 @override 주석에서 class 값 추출
package stml_design

import (
	"golang.org/x/net/html"
)

// walkForOverrides recursively walks the DOM and extracts class values from
// @override comment nodes.
func walkForOverrides(n *html.Node, classes map[string]bool) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.CommentNode && isOverrideComment(c.Data) {
			cls := extractOverrideClass(c.Data)
			if cls != "" {
				classes[cls] = true
			}
		}
		walkForOverrides(c, classes)
	}
}
