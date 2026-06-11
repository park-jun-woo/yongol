//ff:func feature=gen-react type=test control=sequence
//ff:what writeLayoutTSX auth 없음 — data-logout 선언에도 로그아웃 방출 생략 (TM-38 영역) 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteLayoutTSX_LogoutNoAuth(t *testing.T) {
	dir := t.TempDir()
	layout := stml.LayoutSpec{
		Name:      "app",
		File:      "layouts/app.html",
		NavItems:  []stml.NavItem{{Path: "dashboard", Label: "대시보드"}},
		HasOutlet: true,
		Logout:    &stml.LogoutSpec{OperationID: "Logout", Label: "로그아웃"},
	}
	patterns := map[string]string{"dashboard": "/dashboard"}
	if err := writeLayoutTSX(dir, layout, patterns, "", nil); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "layouts", "AppLayout.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	// There is no session to end — the declaration is dead (TM-38 warns)
	// and the emission stays byte-compatible with a logout-less layout.
	assertContains(t, content, "import { Link, Outlet } from 'react-router-dom'")
	assertNotContains(t, content, "useNavigate")
	assertNotContains(t, content, "handleLogout")
	assertNotContains(t, content, "api")
	assertNotContains(t, content, "useAuthStore")
	assertContains(t, content, `<Link to="/dashboard">대시보드</Link>`)
	assertContains(t, content, "<Outlet />")
}
