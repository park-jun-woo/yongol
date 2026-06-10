//ff:func feature=gen-react type=test control=sequence
//ff:what writeLayoutTSX bearer 모드 — 페이지명 nav 치환 + store clear 로그아웃 방출 스냅샷 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteLayoutTSX_LogoutBearer(t *testing.T) {
	dir := t.TempDir()
	layout := stml.LayoutSpec{
		Name: "app",
		File: "layouts/app.html",
		NavItems: []stml.NavItem{
			{Path: "dashboard", Label: "대시보드"},
			{Path: "/legacy", Label: "Legacy"},
		},
		HasOutlet: true,
		Logout:    &stml.LogoutSpec{OperationID: "Logout", Label: "로그아웃"},
	}
	patterns := map[string]string{"dashboard": "/dashboard"}
	if err := writeLayoutTSX(dir, layout, patterns, "bearer"); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "layouts", "AppLayout.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import { Link, Outlet, useNavigate } from 'react-router-dom'")
	assertContains(t, content, "import { useAuthStore } from '@/stores/auth'")
	assertContains(t, content, "import { api } from '@/lib/api'")
	assertContains(t, content, "const navigate = useNavigate()")
	assertContains(t, content, "await api.Logout().catch(() => {})")
	assertContains(t, content, "useAuthStore.getState().clear()")
	assertContains(t, content, "navigate('/login')")
	assertContains(t, content, `<Link to="/dashboard">대시보드</Link>`)
	assertContains(t, content, `<Link to="/legacy">Legacy</Link>`)
	assertContains(t, content, "<button onClick={handleLogout}>로그아웃</button>")
	assertContains(t, content, "<Outlet />")
}
