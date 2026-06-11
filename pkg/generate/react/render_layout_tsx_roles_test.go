//ff:func feature=gen-react type=test control=sequence
//ff:what 레이아웃 TSX role 메뉴 — ROLES 상수·userRole 셀렉터·조건 렌더(단일/복수 role·서브트리 상속) 스냅샷과 role_field 미배선 시 무조건 렌더 검증

package react

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestRenderLayoutTSX_RoleMenu(t *testing.T) {
	items := []sitemapMenuItem{
		{Kind: "page", Label: "대시보드", To: "/dashboard"},
		{Kind: "page", Label: "멤버", To: "/members", Roles: []string{"admin", "manager"}},
		{Kind: "group", Label: "운영", Roles: []string{"admin"}, Children: []sitemapMenuItem{
			{Kind: "page", Label: "감사 로그", To: "/audit"},
			{Kind: "page", Label: "정산", To: "/billing", Roles: []string{"admin"}},
		}},
	}
	layout := stml.LayoutSpec{Name: "app", HasOutlet: true}

	t.Run("wired role_field renders the conditional menu snapshot", func(t *testing.T) {
		got := renderLayoutTSX("AppLayout", layout, nil, "", &sitemapMenu{Items: items, RoleField: "role"})
		want := "import { NavLink, Outlet } from 'react-router-dom'\n" +
			"import { Breadcrumb } from '@/components/ui/Breadcrumb'\n" +
			"import { useAuthStore } from '@/stores/auth'\n" +
			"\n" +
			"const ROLES_admin_manager = ['admin', 'manager']\n" +
			"const ROLES_admin = ['admin']\n" +
			"\n" +
			"export default function AppLayout() {\n" +
			"  const userRole = useAuthStore((s) => s.claims['role'])\n" +
			"\n" +
			"  return (\n" +
			"    <div>\n" +
			"      <nav>\n" +
			"        <ul>\n" +
			"          <li><NavLink to=\"/dashboard\" end>대시보드</NavLink></li>\n" +
			"          {ROLES_admin_manager.includes(userRole) && (\n" +
			"            <li><NavLink to=\"/members\" end>멤버</NavLink></li>\n" +
			"          )}\n" +
			"          {ROLES_admin.includes(userRole) && (\n" +
			"            <li>\n" +
			"              <span>운영</span>\n" +
			"              <ul>\n" +
			"                <li><NavLink to=\"/audit\" end>감사 로그</NavLink></li>\n" +
			"                {ROLES_admin.includes(userRole) && (\n" +
			"                  <li><NavLink to=\"/billing\" end>정산</NavLink></li>\n" +
			"                )}\n" +
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

	t.Run("empty role_field renders everything unconditionally", func(t *testing.T) {
		got := renderLayoutTSX("AppLayout", layout, nil, "", &sitemapMenu{Items: items})
		assertNotContains(t, got, "ROLES_")
		assertNotContains(t, got, "userRole")
		assertNotContains(t, got, "useAuthStore")
		assertContains(t, got, "<li><NavLink to=\"/members\" end>멤버</NavLink></li>")
	})

	t.Run("role_field without role-gated items stays byte-identical", func(t *testing.T) {
		plain := []sitemapMenuItem{{Kind: "page", Label: "대시보드", To: "/dashboard"}}
		with := renderLayoutTSX("AppLayout", layout, nil, "", &sitemapMenu{Items: plain, RoleField: "role"})
		without := renderLayoutTSX("AppLayout", layout, nil, "", &sitemapMenu{Items: plain})
		if with != without {
			t.Errorf("role_field alone must not change the output:\n%s\nvs\n%s", with, without)
		}
		if strings.Contains(with, "userRole") {
			t.Errorf("no role-gated item: userRole must not be emitted")
		}
	})
}
