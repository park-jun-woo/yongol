//ff:func feature=gen-react type=generator control=sequence
//ff:what renderSitemapMenu — sitemap 파생 레이아웃 메뉴 <nav> 블록 방출 (2단 그룹 + role 조건 + 선택적 로그아웃 버튼)

package react

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderSitemapMenu writes the sitemap-derived <nav> block of one layout
// (plans/stml/sitemap Phase003): a <ul> of menu items — 2-level groups
// expanded by default, NavLink active states, external links, icons —
// plus the data-logout button when emitLogout, exactly where the data-nav
// path puts it. Nothing is written when the layout has neither items nor
// a logout button. With menu.RoleField wired (Phase005), data-roles items
// render inside a role condition; a "" RoleField renders everything
// unconditionally (byte-identical pre-Phase005 output).
func renderSitemapMenu(sb *strings.Builder, menu *sitemapMenu, layout stml.LayoutSpec, emitLogout bool) {
	if len(menu.Items) == 0 && !emitLogout {
		return
	}
	rolesActive := menu.RoleField != ""
	sb.WriteString("      <nav>\n")
	if len(menu.Items) > 0 {
		sb.WriteString("        <ul>\n")
		for _, it := range menu.Items {
			renderSitemapItem(sb, it, "          ", rolesActive)
		}
		sb.WriteString("        </ul>\n")
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
