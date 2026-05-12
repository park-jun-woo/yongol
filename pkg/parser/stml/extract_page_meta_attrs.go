//ff:func feature=stml-parse type=parser control=sequence
//ff:what extractPageMetaAttrs — 요소에서 data-layout, data-route 속성을 PageSpec 에 설정

package stml

import "golang.org/x/net/html"

// extractPageMetaAttrs extracts data-layout and data-route attributes
// from the given element node, setting them on page only if not already set.
func extractPageMetaAttrs(n *html.Node, page *PageSpec) {
	if page.Layout == "" {
		if v := getAttr(n, "data-layout"); v != "" {
			page.Layout = v
		}
	}
	if page.Route == "" {
		if v := getAttr(n, "data-route"); v != "" {
			page.Route = v
		}
	}
}
