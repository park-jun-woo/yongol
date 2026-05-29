//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX 인증 가드 + 레이아웃 + flat 혼합 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_Authz_MixedLayoutAndFlat(t *testing.T) {
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
	if err := writeAppTSX(dir, pages, layouts, "", true); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "<Route element={<ProtectedRoute><AppLayout /></ProtectedRoute>}>")
	assertContains(t, content, "<Route element={<AuthLayout />}>")
	assertContains(t, content, `<Route path="/about" element={<ProtectedRoute><About /></ProtectedRoute>} />`)
}
