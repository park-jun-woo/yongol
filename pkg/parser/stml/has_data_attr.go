//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what 요소에 data-* 접두사 속성이 있는지 확인
package stml

import (
	"strings"

	"golang.org/x/net/html"
)

func hasDataAttr(n *html.Node) bool {
	for _, attr := range n.Attr {
		if strings.HasPrefix(attr.Key, "data-") {
			return true
		}
	}
	return false
}
