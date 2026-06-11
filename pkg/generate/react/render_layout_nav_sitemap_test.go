//ff:func feature=gen-react type=test control=sequence
//ff:what renderLayoutNav sitemap 위임 — menu 비nil 시 data-nav 무시 + renderSitemapMenu 경로 검증

package react

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderLayoutNav_SitemapDelegation(t *testing.T) {
	layout := stml.LayoutSpec{
		NavItems: []stml.NavItem{{Path: "/legacy", Label: "Legacy"}},
	}
	menu := &sitemapMenu{Items: []sitemapMenuItem{
		{Kind: "page", Label: "대시보드", To: "/dashboard"},
	}}

	var sb strings.Builder
	renderLayoutNav(&sb, layout, nil, false, menu)
	out := sb.String()
	assertContains(t, out, `<NavLink to="/dashboard" end>대시보드</NavLink>`)
	assertNotContains(t, out, "Legacy")
	assertNotContains(t, out, "<Link ")
}
