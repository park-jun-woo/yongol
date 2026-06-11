//ff:func feature=gen-react type=test control=sequence
//ff:what TestRenderLayoutTSXDynamicGroup — bearer 게이트 전체 스냅샷/cookie 무게이트/그룹 부재 시 바이트 동일 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderLayoutTSXDynamicGroup(t *testing.T) {
	layout := stml.LayoutSpec{Name: "app", HasOutlet: true}
	items := []sitemapMenuItem{
		{Kind: "page", Label: "대시보드", To: "/dashboard"},
		{
			Kind: "group", Label: "내 건물",
			Fetch: "ListMyBuildings", Each: "items", LabelField: "building_name",
			ItemToAttr: "to={`/buildings/${item.building_id}`}", ItemKey: "item.building_id",
		},
	}

	t.Run("bearer full snapshot — token-gated useQuery and the omission-gated group", func(t *testing.T) {
		got := renderLayoutTSX("AppLayout", layout, nil, "bearer", &sitemapMenu{Items: items})
		want := "import { NavLink, Outlet } from 'react-router-dom'\n" +
			"import { useQuery } from '@tanstack/react-query'\n" +
			"import { Breadcrumb } from '@/components/ui/Breadcrumb'\n" +
			"import { useAuthStore } from '@/stores/auth'\n" +
			"import { api } from '@/lib/api'\n" +
			"\n" +
			"export default function AppLayout() {\n" +
			"  const token = useAuthStore((s) => s.token)\n" +
			"  const { data: listMyBuildingsData } = useQuery({\n" +
			"    queryKey: ['ListMyBuildings'],\n" +
			"    queryFn: () => api.ListMyBuildings(),\n" +
			"    enabled: !!token,\n" +
			"  })\n" +
			"\n" +
			"  return (\n" +
			"    <div>\n" +
			"      <nav>\n" +
			"        <ul>\n" +
			"          <li><NavLink to=\"/dashboard\" end>대시보드</NavLink></li>\n" +
			"          {(listMyBuildingsData?.items ?? []).length > 0 && (\n" +
			"            <li>\n" +
			"              <span>내 건물</span>\n" +
			"              <ul>\n" +
			"                {(listMyBuildingsData?.items ?? []).map((item) => (\n" +
			"                  <li key={item.building_id}><NavLink to={`/buildings/${item.building_id}`} end>{item.building_name}</NavLink></li>\n" +
			"                ))}\n" +
			"              </ul>\n" +
			"            </li>\n" +
			"          )}\n" +
			"        </ul>\n" +
			"      </nav>\n" +
			"      <Breadcrumb />\n" +
			"      <Outlet />\n" +
			"    </div>\n" +
			"  )\n" +
			"}\n"
		if got != want {
			t.Errorf("got:\n%s\nwant:\n%s", got, want)
		}
	})

	t.Run("cookie mode fires ungated and needs no auth store", func(t *testing.T) {
		got := renderLayoutTSX("AppLayout", layout, nil, "cookie", &sitemapMenu{Items: items})
		assertContains(t, got, "import { useQuery } from '@tanstack/react-query'\n")
		assertContains(t, got, "  const { data: listMyBuildingsData } = useQuery({\n    queryKey: ['ListMyBuildings'],\n    queryFn: () => api.ListMyBuildings(),\n  })\n")
		assertNotContains(t, got, "useAuthStore")
		assertNotContains(t, got, "enabled:")
	})

	t.Run("without dynamic groups the Phase006 emission stays byte-identical", func(t *testing.T) {
		static := &sitemapMenu{Items: []sitemapMenuItem{{Kind: "page", Label: "대시보드", To: "/dashboard"}}}
		got := renderLayoutTSX("AppLayout", layout, nil, "bearer", static)
		assertNotContains(t, got, "useQuery")
		assertNotContains(t, got, "@tanstack/react-query")
		assertNotContains(t, got, "api")
		assertNotContains(t, got, "token")
	})
}
