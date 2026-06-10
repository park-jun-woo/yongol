//ff:func feature=gen-react type=generator control=iteration dimension=1
//ff:what renderLayoutNav — 레이아웃 <nav> 블록 방출 (data-nav Link 들 + 선택적 로그아웃 버튼)

package react

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// renderLayoutNav writes the layout's <nav> block: one <Link> per
// data-nav entry (resolved through navLinkPath) and, when emitLogout, the
// data-logout button wired to handleLogout. Nothing is written when the
// layout has neither — the <nav> wrapper exists only to host content.
func renderLayoutNav(sb *strings.Builder, layout stml.LayoutSpec, routePatterns map[string]string, emitLogout bool) {
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
