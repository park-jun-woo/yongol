//ff:func feature=gen-react type=test control=sequence
//ff:what renderLayoutTSX — DynamicCrumb 시 라벨 state·pathname 리셋·label prop·Outlet context 배선 / 미배선 레이아웃 바이트 동일 검증

package react

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderLayoutTSXDynamicCrumb(t *testing.T) {
	items := []sitemapMenuItem{{Kind: "page", Label: "대시보드", To: "/dashboard"}}
	layout := stml.LayoutSpec{Name: "app", HasOutlet: true}

	t.Run("DynamicCrumb wires state, reset, label prop and Outlet context", func(t *testing.T) {
		got := renderLayoutTSX("AppLayout", layout, nil, "", &sitemapMenu{Items: items, DynamicCrumb: true})
		assertContains(t, got, "import { useEffect, useState } from 'react'")
		assertContains(t, got, "import { NavLink, Outlet, useLocation } from 'react-router-dom'")
		assertContains(t, got, "  const { pathname } = useLocation()\n")
		assertContains(t, got, "  const [crumbLabel, setCrumbLabel] = useState<string | null>(null)\n")
		// pathname-change reset — the stale-label guard
		assertContains(t, got, "  useEffect(() => {\n    setCrumbLabel(null)\n  }, [pathname])\n")
		assertContains(t, got, "      <Breadcrumb label={crumbLabel} />\n      <Outlet context={{ setCrumbLabel }} />\n")
		assertNotContains(t, got, "<Outlet />")
	})

	t.Run("without DynamicCrumb the Phase005 emission stays byte-identical", func(t *testing.T) {
		got := renderLayoutTSX("AppLayout", layout, nil, "", &sitemapMenu{Items: items})
		assertNotContains(t, got, "setCrumbLabel")
		assertNotContains(t, got, "useState")
		assertContains(t, got, "      <Breadcrumb />\n      <Outlet />\n")
		if strings.Contains(got, "from 'react'") {
			t.Errorf("react import must not appear without DynamicCrumb:\n%s", got)
		}
	})

	t.Run("DynamicCrumb without an Outlet wires nothing", func(t *testing.T) {
		got := renderLayoutTSX("BareLayout", stml.LayoutSpec{Name: "bare"}, nil, "", &sitemapMenu{Items: items, DynamicCrumb: true})
		assertNotContains(t, got, "setCrumbLabel")
		assertNotContains(t, got, "Breadcrumb")
	})
}
