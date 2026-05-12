//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX 레이아웃 없이 flat 라우트 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_NoLayoutNoDefault_FlatRoutes(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html"},
		{Name: "dashboard", FileName: "dashboard.html"},
	}
	if err := writeAppTSX(dir, pages, nil, "", false); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, `      <Route path="/dashboard" element={<Dashboard />} />`)
	assertContains(t, content, `      <Route path="/login" element={<Login />} />`)
	assertNotContains(t, content, "Layout")
}
