//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what 요소의 첫 번째 비어있지 않은 직접 텍스트 자식 추출
package stml

import "golang.org/x/net/html"

// directText extracts the first non-empty direct text child of an element.
func directText(n *html.Node) string {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if text := extractNonEmptyText(c); text != "" {
			return text
		}
	}
	return ""
}
