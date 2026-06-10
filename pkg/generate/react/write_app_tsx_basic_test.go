//ff:func feature=gen-react type=test control=sequence
//ff:what writeAppTSX 기본 페이지 라우트 생성 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteAppTSX_BasicPages(t *testing.T) {
	dir := t.TempDir()
	pages := []stml.PageSpec{
		{Name: "login", FileName: "login.html"},
		{Name: "register", FileName: "register.html"},
		{Name: "dashboard", FileName: "dashboard.html"},
	}
	if err := writeAppTSX(dir, pages, nil, "", nil); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "App.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import Login from './pages/login'")
	assertContains(t, content, "import Register from './pages/register'")
	assertContains(t, content, "import Dashboard from './pages/dashboard'")
	assertContains(t, content, `<Route path="/login" element={<Login />} />`)
	assertContains(t, content, `<Route path="/register" element={<Register />} />`)
	assertContains(t, content, `<Route path="/dashboard" element={<Dashboard />} />`)
}
