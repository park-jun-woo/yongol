//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX 레이아웃 그룹 내 페이지별 가드 검증 (래퍼 비가드)

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_Authz_LayoutGrouping(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html", Layout: "auth"},
		{Name: "register", FileName: "register.html", Layout: "auth"},
		{Name: "workflows", FileName: "workflows.html", Layout: "app"},
		{Name: "dashboard", FileName: "dashboard.html", Layout: "app"},
	}
	layouts := []stml.LayoutSpec{
		{Name: "app", HasOutlet: true, NavItems: []stml.NavItem{{Path: "/workflows", Label: "Workflows"}}},
		{Name: "auth", HasOutlet: true},
	}
	protected := map[string]bool{"workflows.html": true, "dashboard.html": true}
	if err := writeAppTSX(dir, pages, layouts, "", protected, ""); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import ProtectedRoute from './components/ProtectedRoute'")
	// layout wrappers stay unguarded — each protected page wraps itself
	assertContains(t, content, "<Route element={<AppLayout />}>")
	assertContains(t, content, "<Route element={<AuthLayout />}>")
	assertNotContains(t, content, "<ProtectedRoute><AppLayout /></ProtectedRoute>")
	assertNotContains(t, content, "<ProtectedRoute><AuthLayout /></ProtectedRoute>")
	assertContains(t, content, `        <Route path="/workflows" element={<ProtectedRoute><Workflows /></ProtectedRoute>} />`)
	assertContains(t, content, `        <Route path="/dashboard" element={<ProtectedRoute><Dashboard /></ProtectedRoute>} />`)
	assertContains(t, content, `        <Route path="/login" element={<Login />} />`)
	assertContains(t, content, `        <Route path="/register" element={<Register />} />`)
}
