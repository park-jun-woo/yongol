//ff:func feature=gen-react type=generator control=sequence
//ff:what renderLayoutOutlet — 레이아웃의 Breadcrumb+Outlet 방출 (sitemap 시 Breadcrumb, 동적 crumb 시 label prop + Outlet context)

package react

import "strings"

// renderLayoutOutlet writes the layout's Breadcrumb + Outlet pair
// (plans/stml/sitemap Phase004/006). hasMenu (sitemap present) places the
// shared <Breadcrumb> above the Outlet; dynamicCrumb additionally feeds
// the crumb-label state through the label prop and hands setCrumbLabel
// down via the react-router Outlet context — the official layout→page
// channel, so no dependency is added. Without dynamicCrumb the Phase004
// bytes stay identical.
func renderLayoutOutlet(sb *strings.Builder, hasMenu, dynamicCrumb bool) {
	if hasMenu && dynamicCrumb {
		sb.WriteString("      <Breadcrumb label={crumbLabel} />\n")
	} else if hasMenu {
		sb.WriteString("      <Breadcrumb />\n")
	}
	if dynamicCrumb {
		sb.WriteString("      <Outlet context={{ setCrumbLabel }} />\n")
	} else {
		sb.WriteString("      <Outlet />\n")
	}
}
