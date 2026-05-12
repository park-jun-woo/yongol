//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX defaultLayout 적용 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_DefaultLayout(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html", Layout: "auth"},
		{Name: "workflows", FileName: "workflows.html"},
		{Name: "dashboard", FileName: "dashboard.html"},
	}
	layouts := []stml.LayoutSpec{
		{Name: "app", HasOutlet: true},
		{Name: "auth", HasOutlet: true},
	}
	if err := writeAppTSX(dir, pages, layouts, "app", false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "<Route element={<AppLayout />}>")
	assertContains(t, content, `        <Route path="/workflows" element={<Workflows />} />`)
	assertContains(t, content, `        <Route path="/dashboard" element={<Dashboard />} />`)
	assertContains(t, content, "<Route element={<AuthLayout />}>")
	assertContains(t, content, `        <Route path="/login" element={<Login />} />`)
}
