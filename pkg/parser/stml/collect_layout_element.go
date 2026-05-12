//ff:func feature=stml-parse type=parser control=sequence
//ff:what collectLayoutElement — 레이아웃 HTML 요소에서 data-nav 링크 및 data-outlet 슬롯 수집

package stml

import "golang.org/x/net/html"

// collectLayoutElement checks a single element node for data-nav on <a>
// and data-outlet on <slot>.
func collectLayoutElement(n *html.Node, layout *LayoutSpec) {
	if n.Data == "a" {
		if nav := getAttr(n, "data-nav"); nav != "" {
			label := extractAllText(n)
			layout.NavItems = append(layout.NavItems, NavItem{
				Path:  nav,
				Label: label,
			})
		}
	}
	if n.Data == "slot" && hasAttr(n, "data-outlet") {
		layout.HasOutlet = true
	}
}
