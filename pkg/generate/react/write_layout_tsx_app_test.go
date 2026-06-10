//ff:func feature=gen-react type=test control=sequence
//ff:what writeLayoutTSX App 레이아웃 TSX 생성 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteLayoutTSX_AppLayout(t *testing.T) {
	dir := t.TempDir()
	layout := stml.LayoutSpec{
		Name: "app",
		File: "layouts/app.html",
		NavItems: []stml.NavItem{
			{Path: "/workflows", Label: "Workflows"},
			{Path: "/dashboard", Label: "Dashboard"},
		},
		HasOutlet: true,
	}
	if err := writeLayoutTSX(dir, layout, nil, ""); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "layouts", "AppLayout.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import { Link, Outlet } from 'react-router-dom'")
	assertContains(t, content, "export default function AppLayout()")
	assertContains(t, content, `<Link to="/workflows">Workflows</Link>`)
	assertContains(t, content, `<Link to="/dashboard">Dashboard</Link>`)
	assertContains(t, content, "<Outlet />")
	assertContains(t, content, "<nav>")
}
