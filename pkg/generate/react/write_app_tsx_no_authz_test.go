//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX 인증 부재 시 ProtectedRoute 미포함 + 인덱스/catch-all 방출 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_NoAuthz_NoProtectedRoute(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html", Layout: "auth"},
		{Name: "workflows", FileName: "workflows.html", Layout: "app"},
	}
	layouts := []stml.LayoutSpec{
		{Name: "app", HasOutlet: true},
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

	assertNotContains(t, content, "ProtectedRoute")
	assertContains(t, content, "<Route element={<AppLayout />}>")
	assertContains(t, content, "<Route element={<AuthLayout />}>")
	// index/catch-all are emitted regardless of auth (BUG-111 (5) UX gap)
	assertContains(t, content, `<Route path="/" element={<Navigate to="/login" replace />} />`)
	assertContains(t, content, `<Route path="*" element={<Navigate to="/" replace />} />`)
}
