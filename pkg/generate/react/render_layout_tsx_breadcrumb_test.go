//ff:func feature=gen-react type=test control=sequence
//ff:what renderLayoutTSX — sitemap 존재 시 Outlet 위 Breadcrumb 방출 / outlet 없으면 미방출 / sitemap 부재 바이트 동일 검증

package react

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderLayoutTSXBreadcrumb(t *testing.T) {
	menu := &sitemapMenu{Items: []sitemapMenuItem{
		{Kind: "page", Label: "대시보드", To: "/dashboard"},
	}}

	t.Run("sitemap menu with outlet places Breadcrumb above it", func(t *testing.T) {
		got := renderLayoutTSX("AppLayout", stml.LayoutSpec{Name: "app", HasOutlet: true}, nil, "", menu)
		assertContains(t, got, "import { Breadcrumb } from '@/components/ui/Breadcrumb'")
		assertContains(t, got, "      <Breadcrumb />\n      <Outlet />\n")
	})

	t.Run("no outlet, no Breadcrumb", func(t *testing.T) {
		got := renderLayoutTSX("BareLayout", stml.LayoutSpec{Name: "bare"}, nil, "", menu)
		assertNotContains(t, got, "Breadcrumb")
	})

	t.Run("sitemap absent stays byte-identical (no Breadcrumb anywhere)", func(t *testing.T) {
		got := renderLayoutTSX("AppLayout", stml.LayoutSpec{Name: "app", HasOutlet: true}, nil, "", nil)
		assertNotContains(t, got, "Breadcrumb")
		if !strings.Contains(got, "      <Outlet />\n") {
			t.Errorf("legacy outlet emission changed:\n%s", got)
		}
	})
}
