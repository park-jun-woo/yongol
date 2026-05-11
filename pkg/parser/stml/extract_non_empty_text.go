//ff:func feature=stml-parse type=util control=sequence
//ff:what 텍스트 노드에서 비어있지 않은 텍스트를 추출
package stml

import (
	"strings"

	"golang.org/x/net/html"
)

func extractNonEmptyText(c *html.Node) string {
	if c.Type != html.TextNode {
		return ""
	}
	return strings.TrimSpace(c.Data)
}
