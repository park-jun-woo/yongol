//ff:func feature=gen-react type=test control=sequence
//ff:what renderSitemapMenu — 전체 <nav> 스냅샷 (그룹/직속/외부 링크/로그아웃 버튼/빈 메뉴 생략) 검증

package react

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderSitemapMenu(t *testing.T) {
	items := []sitemapMenuItem{
		{Kind: "page", Label: "대시보드", To: "/dashboard"},
		{Kind: "group", Label: "건물 관리", Children: []sitemapMenuItem{
			{Kind: "page", Label: "건물 목록", To: "/buildings", Prefixes: []string{"/buildings/"}},
		}},
		{Kind: "external", Label: "매뉴얼", Href: "https://docs.example.com"},
	}

	t.Run("full nav snapshot with logout", func(t *testing.T) {
		var sb strings.Builder
		layout := stml.LayoutSpec{Logout: &stml.LogoutSpec{Label: "로그아웃"}}
		renderSitemapMenu(&sb, &sitemapMenu{Items: items}, layout, true)
		want := "      <nav>\n" +
			"        <ul>\n" +
			"          <li><NavLink to=\"/dashboard\" end>대시보드</NavLink></li>\n" +
			"          <li>\n" +
			"            <span>건물 관리</span>\n" +
			"            <ul>\n" +
			"              <li><NavLink to=\"/buildings\" end className={({ isActive }) => (isActive || pathname.startsWith('/buildings/') ? 'active' : undefined)}>건물 목록</NavLink></li>\n" +
			"            </ul>\n" +
			"          </li>\n" +
			"          <li><a href=\"https://docs.example.com\" target=\"_blank\" rel=\"noopener noreferrer\">매뉴얼</a></li>\n" +
			"        </ul>\n" +
			"        <button onClick={handleLogout}>로그아웃</button>\n" +
			"      </nav>\n"
		if sb.String() != want {
			t.Errorf("got:\n%s\nwant:\n%s", sb.String(), want)
		}
	})

	t.Run("nothing without items or logout", func(t *testing.T) {
		var sb strings.Builder
		renderSitemapMenu(&sb, &sitemapMenu{}, stml.LayoutSpec{}, false)
		if sb.Len() != 0 {
			t.Errorf("expected empty output, got %q", sb.String())
		}
	})

	t.Run("logout-only keeps the nav wrapper", func(t *testing.T) {
		var sb strings.Builder
		renderSitemapMenu(&sb, &sitemapMenu{}, stml.LayoutSpec{Logout: &stml.LogoutSpec{}}, true)
		out := sb.String()
		assertContains(t, out, "<nav>")
		assertContains(t, out, "<button onClick={handleLogout}>Logout</button>")
		assertNotContains(t, out, "<ul>")
	})
}
