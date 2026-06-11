//ff:func feature=gen-react type=generator control=sequence
//ff:what renderSitemapItem — 메뉴 항목 방출: data-roles 항목은 role 조건 블록으로 감싸고 본문은 위임

package react

import (
	"fmt"
	"strings"
)

// renderSitemapItem writes one menu item. An item carrying a data-roles
// allowlist (and an active role wiring — plans/stml/sitemap Phase005) is
// wrapped in a conditional render block:
//
//	{ROLES_admin.includes(userRole) && (
//	  <li>…</li>
//	)}
//
// userRole is the layout-level claims[role_field] selector and ROLES_* the
// module-level allowlist constant. A signed-out user has no claim →
// includes(undefined) is false → the item is hidden. Children render
// inside the parent's <li>, so an ancestor's condition ANDs over the
// subtree by nesting alone. Items without roles (or without wiring)
// delegate straight to renderSitemapItemBody — byte-identical to the
// pre-Phase005 output.
func renderSitemapItem(sb *strings.Builder, item sitemapMenuItem, indent string, rolesActive bool) {
	if rolesActive && len(item.Roles) > 0 {
		fmt.Fprintf(sb, "%s{%s.includes(userRole) && (\n", indent, rolesConstName(item.Roles))
		renderSitemapItemBody(sb, item, indent+"  ", rolesActive)
		fmt.Fprintf(sb, "%s)}\n", indent)
		return
	}
	renderSitemapItemBody(sb, item, indent, rolesActive)
}
