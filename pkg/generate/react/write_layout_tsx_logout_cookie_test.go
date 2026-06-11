//ff:func feature=gen-react type=test control=sequence
//ff:what writeLayoutTSX cookie 모드 — 서버 op 호출 로그아웃 방출, store import 부재 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteLayoutTSX_LogoutCookie(t *testing.T) {
	dir := t.TempDir()
	layout := stml.LayoutSpec{
		Name:      "app",
		File:      "layouts/app.html",
		HasOutlet: true,
		Logout:    &stml.LogoutSpec{OperationID: "Logout", Label: "로그아웃"},
	}
	if err := writeLayoutTSX(dir, layout, nil, "cookie", nil); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "layouts", "AppLayout.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import { Outlet, useNavigate } from 'react-router-dom'")
	assertContains(t, content, "import { api } from '@/lib/api'")
	assertNotContains(t, content, "useAuthStore")
	assertContains(t, content, "await api.Logout().catch(() => {})")
	assertContains(t, content, "navigate('/login')")
	// no nav items: the <nav> exists solely to host the logout button.
	assertContains(t, content, "<nav>")
	assertContains(t, content, "<button onClick={handleLogout}>로그아웃</button>")
	assertNotContains(t, content, "<Link")
}
