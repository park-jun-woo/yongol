//ff:func feature=stml-parse type=parser control=sequence
//ff:what collectLayoutElement — 레이아웃 HTML 요소에서 data-nav 링크, data-outlet 슬롯, data-logout 마커 수집

package stml

import "golang.org/x/net/html"

// collectLayoutElement checks a single element node for data-nav on <a>,
// data-outlet on <slot>, and data-logout on any clickable element
// (page-flow Phase010). The parser keeps the data-nav value verbatim —
// resolving page-name references is validate/generate work — and the
// first data-logout occurrence wins (a layout has one session to end).
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
	if layout.Logout == nil && hasAttr(n, "data-logout") {
		layout.Logout = &LogoutSpec{
			OperationID: getAttr(n, "data-logout"),
			Label:       extractAllText(n),
		}
	}
}
