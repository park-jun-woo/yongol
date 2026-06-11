//ff:func feature=gen-react type=test control=sequence
//ff:what renderSitemapItem — 리프 한 줄/자식 있는 항목의 중첩 <ul> 블록 방출 검증

package react

import (
	"strings"
	"testing"
)

func TestRenderSitemapItem(t *testing.T) {
	t.Run("leaf renders on one line", func(t *testing.T) {
		var sb strings.Builder
		renderSitemapItem(&sb, sitemapMenuItem{Kind: "page", Label: "대시보드", To: "/dashboard"}, "  ", false)
		want := "  <li><NavLink to=\"/dashboard\" end>대시보드</NavLink></li>\n"
		if sb.String() != want {
			t.Errorf("got %q, want %q", sb.String(), want)
		}
	})

	t.Run("item with children expands a nested ul", func(t *testing.T) {
		var sb strings.Builder
		renderSitemapItem(&sb, sitemapMenuItem{
			Kind: "group", Label: "건물 관리",
			Children: []sitemapMenuItem{{Kind: "page", Label: "건물 목록", To: "/buildings"}},
		}, "  ", false)
		want := "  <li>\n" +
			"    <span>건물 관리</span>\n" +
			"    <ul>\n" +
			"      <li><NavLink to=\"/buildings\" end>건물 목록</NavLink></li>\n" +
			"    </ul>\n" +
			"  </li>\n"
		if sb.String() != want {
			t.Errorf("got:\n%s\nwant:\n%s", sb.String(), want)
		}
	})
}
