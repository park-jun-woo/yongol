//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX 전부 보호 페이지 — flat 라우트 페이지별 가드 + /login 인덱스 폴백 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_Authz_FlatRoutes(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "about", FileName: "about.html"},
		{Name: "settings", FileName: "settings.html"},
	}
	protected := map[string]bool{"about.html": true, "settings.html": true}
	if err := writeAppTSX(dir, pages, nil, "", protected, ""); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import ProtectedRoute from './components/ProtectedRoute'")
	assertContains(t, content, `<Route path="/about" element={<ProtectedRoute><About /></ProtectedRoute>} />`)
	assertContains(t, content, `<Route path="/settings" element={<ProtectedRoute><Settings /></ProtectedRoute>} />`)
	// every page is protected → the "/" index falls back to /login
	assertContains(t, content, `<Route path="/" element={<Navigate to="/login" replace />} />`)
	assertContains(t, content, `<Route path="*" element={<Navigate to="/" replace />} />`)
}
