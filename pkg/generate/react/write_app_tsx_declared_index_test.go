//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX "/" 인덱스 — manifest frontend.index 선언 페이지로 redirect 방출 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_IndexRedirect_DeclaredIndex(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "forgot-password", FileName: "forgot-password.html"},
		{Name: "dashboard", FileName: "dashboard.html"},
	}
	// dashboard is protected, but a declared index may be protected —
	// <ProtectedRoute> bounces unauthenticated visits to /login after the
	// redirect (page-flow Phase009). Without the declaration the fallback
	// would pick /forgot-password (BUG-114 (3)).
	protected := map[string]bool{"dashboard.html": true}
	if err := writeAppTSX(dir, pages, nil, "", protected, "dashboard", nil); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, `<Route path="/" element={<Navigate to="/dashboard" replace />} />`)
	assertContains(t, content, `<Route path="*" element={<Navigate to="/" replace />} />`)
}
