//ff:func feature=gen-react type=generator control=selection
//ff:what renderSitemapEntry — 메뉴 항목 한 개의 인라인 마크업 (NavLink end + 조상 prefix className / 외부 <a> / 그룹 <span>, 아이콘 포함)

package react

import (
	"fmt"
	"strings"
)

// renderSitemapEntry renders the inline markup of one menu item: an
// internal page becomes <NavLink end> (exact active matching; when the
// item has menu-hidden descendants a className callback extends the match
// with pathname.startsWith over their static route prefixes — the DESIGN
// §4.4 nearest-rendered-ancestor highlight), an external link becomes
// <a target="_blank" rel="noopener noreferrer">, and a group label a
// non-clickable <span>. A data-icon renders as its lucide-react component
// before the label.
func renderSitemapEntry(item sitemapMenuItem) string {
	label := item.Label
	if item.Icon != "" {
		label = fmt.Sprintf("<%s /> %s", item.Icon, label)
	}
	switch item.Kind {
	case "page":
		if len(item.Prefixes) == 0 {
			return fmt.Sprintf("<NavLink to=\"%s\" end>%s</NavLink>", item.To, label)
		}
		conds := make([]string, 0, len(item.Prefixes))
		for _, p := range item.Prefixes {
			conds = append(conds, fmt.Sprintf("pathname.startsWith('%s')", p))
		}
		return fmt.Sprintf("<NavLink to=\"%s\" end className={({ isActive }) => (isActive || %s ? 'active' : undefined)}>%s</NavLink>",
			item.To, strings.Join(conds, " || "), label)
	case "external":
		return fmt.Sprintf("<a href=\"%s\" target=\"_blank\" rel=\"noopener noreferrer\">%s</a>", item.Href, label)
	default:
		return fmt.Sprintf("<span>%s</span>", label)
	}
}
