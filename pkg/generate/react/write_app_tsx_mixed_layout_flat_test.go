//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX 레이아웃 + flat 혼합 라우트 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_MixedLayoutAndFlat(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html", Layout: "auth"},
		{Name: "about", FileName: "about.html"},
		{Name: "workflows", FileName: "workflows.html", Layout: "app"},
	}
	layouts := []stml.LayoutSpec{
		{Name: "app", HasOutlet: true},
		{Name: "auth", HasOutlet: true},
	}
	if err := writeAppTSX(dir, pages, layouts, "", false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "<Route element={<AppLayout />}>")
	assertContains(t, content, "<Route element={<AuthLayout />}>")
	assertContains(t, content, `      <Route path="/about" element={<About />} />`)
	assertContains(t, content, "import AppLayout from './layouts/AppLayout'")
	assertContains(t, content, "import AuthLayout from './layouts/AuthLayout'")
}
