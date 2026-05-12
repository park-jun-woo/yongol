//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what collectText — HTML 노드 트리에서 텍스트 노드 내용을 재귀적으로 Builder 에 수집

package stml

import (
	"strings"

	"golang.org/x/net/html"
)

// collectText recursively collects text node content.
func collectText(n *html.Node, sb *strings.Builder) {
	if n.Type == html.TextNode {
		sb.WriteString(n.Data)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		collectText(c, sb)
	}
}
