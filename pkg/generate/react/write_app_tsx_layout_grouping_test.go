//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX 레이아웃 그룹핑 라우트 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_LayoutGrouping(t *testing.T) {
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
	if err := writeAppTSX(dir, pages, layouts, "", nil, "", nil); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import AppLayout from './layouts/AppLayout'")
	assertContains(t, content, "import AuthLayout from './layouts/AuthLayout'")
	assertContains(t, content, "<Route element={<AppLayout />}>")
	assertContains(t, content, "<Route element={<AuthLayout />}>")
	assertContains(t, content, `        <Route path="/workflows" element={<Workflows />} />`)
	assertContains(t, content, `        <Route path="/dashboard" element={<Dashboard />} />`)
	assertContains(t, content, `        <Route path="/login" element={<Login />} />`)
	assertContains(t, content, `        <Route path="/register" element={<Register />} />`)
}
