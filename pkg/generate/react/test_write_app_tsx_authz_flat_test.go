//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX 인증 가드 + flat 라우트 검증

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
	if err := writeAppTSX(dir, pages, nil, "", true); err != nil {
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
}
