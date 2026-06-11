//ff:func feature=gen-react type=generator control=iteration dimension=1
//ff:what renderLayoutNav — 레이아웃 <nav> 블록 방출 (sitemap 존재 시 sitemap 메뉴 위임, 부재 시 data-nav Link 들 + 선택적 로그아웃 버튼)

package react

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderLayoutNav writes the layout's <nav> block. With a sitemap-derived
// menu (menu != nil — frontend/sitemap.html exists, plans/stml/sitemap
// Phase003) emission delegates to renderSitemapMenu and data-nav is
// ignored (TM-44 rejects its coexistence). Without one (menu == nil) the
// legacy path stays byte-identical: one <Link> per data-nav entry
// (resolved through navLinkPath) and, when emitLogout, the data-logout
// button wired to handleLogout. Nothing is written when the layout has
// neither — the <nav> wrapper exists only to host content.
func renderLayoutNav(sb *strings.Builder, layout stml.LayoutSpec, routePatterns map[string]string, emitLogout bool, menu *sitemapMenu) {
	if menu != nil {
		renderSitemapMenu(sb, menu, layout, emitLogout)
		return
	}
	if len(layout.NavItems) == 0 && !emitLogout {
		return
	}
	sb.WriteString("      <nav>\n")
	for _, item := range layout.NavItems {
		fmt.Fprintf(sb, "        <Link to=\"%s\">%s</Link>\n", navLinkPath(item.Path, routePatterns), item.Label)
	}
	if emitLogout {
		label := layout.Logout.Label
		if label == "" {
			label = "Logout"
		}
		fmt.Fprintf(sb, "        <button onClick={handleLogout}>%s</button>\n", label)
	}
	sb.WriteString("      </nav>\n")
}
