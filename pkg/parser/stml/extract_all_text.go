//ff:func feature=stml-parse type=util control=sequence
//ff:what extractAllText — HTML 요소와 하위 노드의 텍스트 내용을 재귀 수집하여 반환

package stml

import (
	"strings"

	"golang.org/x/net/html"
)

// extractAllText collects all text content from an element and its descendants.
func extractAllText(n *html.Node) string {
	var sb strings.Builder
	collectText(n, &sb)
	return strings.TrimSpace(sb.String())
}
