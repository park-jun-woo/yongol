//ff:func feature=gen-react type=test control=sequence
//ff:what renderSitemapEntry — NavLink end 정확 매칭/조상 prefix className/외부 링크/그룹 span/아이콘 방출 검증

package react

import "testing"

func TestRenderSitemapEntry(t *testing.T) {
	t.Run("page without prefixes", func(t *testing.T) {
		got := renderSitemapEntry(sitemapMenuItem{Kind: "page", Label: "대시보드", To: "/dashboard"})
		want := `<NavLink to="/dashboard" end>대시보드</NavLink>`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("page with ancestor-highlight prefixes", func(t *testing.T) {
		got := renderSitemapEntry(sitemapMenuItem{
			Kind: "page", Label: "건물 목록", To: "/buildings",
			Prefixes: []string{"/buildings/", "/building-documents"},
		})
		want := `<NavLink to="/buildings" end className={({ isActive }) => (isActive || pathname.startsWith('/buildings/') || pathname.startsWith('/building-documents') ? 'active' : undefined)}>건물 목록</NavLink>`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("external link", func(t *testing.T) {
		got := renderSitemapEntry(sitemapMenuItem{Kind: "external", Label: "매뉴얼", Href: "https://docs.example.com"})
		want := `<a href="https://docs.example.com" target="_blank" rel="noopener noreferrer">매뉴얼</a>`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("group label", func(t *testing.T) {
		got := renderSitemapEntry(sitemapMenuItem{Kind: "group", Label: "건물 관리"})
		if want := `<span>건물 관리</span>`; got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("icon precedes the label", func(t *testing.T) {
		got := renderSitemapEntry(sitemapMenuItem{Kind: "page", Label: "대시보드", To: "/dashboard", Icon: "LayoutDashboard"})
		want := `<NavLink to="/dashboard" end><LayoutDashboard /> 대시보드</NavLink>`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
