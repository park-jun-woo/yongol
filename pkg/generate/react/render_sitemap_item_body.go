//ff:func feature=gen-react type=generator control=iteration dimension=1
//ff:what renderSitemapItemBody — 메뉴 항목 <li> 방출 (자식 있으면 중첩 <ul> 재귀, 없으면 한 줄)

package react

import (
	"fmt"
	"strings"
)

// renderSitemapItemBody writes one menu item <li>: a single line for leaf
// items, an expanded block with a nested <ul> when the item has rendered
// children (group → item, the 2-level menu of DESIGN §4.2 — groups are
// always expanded, collapse is out of scope). A dynamic menu group
// (Phase007 — item.Fetch set) delegates to renderSitemapDynamicGroup,
// whose fetched items take the place of static children. indent is the
// <li>'s own indentation; nesting adds four spaces (ul + li) per level.
// Children recurse through renderSitemapItem so a role-gated child wraps
// itself (Phase005).
func renderSitemapItemBody(sb *strings.Builder, item sitemapMenuItem, indent string, rolesActive bool) {
	if item.Fetch != "" {
		renderSitemapDynamicGroup(sb, item, indent, rolesActive)
		return
	}
	entry := renderSitemapEntry(item)
	if len(item.Children) == 0 {
		fmt.Fprintf(sb, "%s<li>%s</li>\n", indent, entry)
		return
	}
	fmt.Fprintf(sb, "%s<li>\n", indent)
	fmt.Fprintf(sb, "%s  %s\n", indent, entry)
	fmt.Fprintf(sb, "%s  <ul>\n", indent)
	for _, c := range item.Children {
		renderSitemapItem(sb, c, indent+"    ", rolesActive)
	}
	fmt.Fprintf(sb, "%s  </ul>\n", indent)
	fmt.Fprintf(sb, "%s</li>\n", indent)
}
