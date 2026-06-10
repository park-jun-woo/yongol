//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX 공개/보호 혼재 — 페이지별 가드 + 레이아웃 래퍼 비가드 검증

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
	protected := map[string]bool{"about.html": true, "workflows.html": true}
	if err := writeAppTSX(dir, pages, layouts, "", protected); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	// layout wrappers are never guarded — protection is per page route
	assertContains(t, content, "<Route element={<AppLayout />}>")
	assertContains(t, content, "<Route element={<AuthLayout />}>")
	assertNotContains(t, content, "<ProtectedRoute><AppLayout /></ProtectedRoute>")
	assertContains(t, content, `        <Route path="/workflows" element={<ProtectedRoute><Workflows /></ProtectedRoute>} />`)
	assertContains(t, content, `        <Route path="/login" element={<Login />} />`)
	assertContains(t, content, `      <Route path="/about" element={<ProtectedRoute><About /></ProtectedRoute>} />`)
	// first public page in file-name order: about(protected) → login
	assertContains(t, content, `<Route path="/" element={<Navigate to="/login" replace />} />`)
}
