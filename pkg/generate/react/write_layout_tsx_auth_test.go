//ff:func feature=gen-react type=test control=sequence
//ff:what writeLayoutTSX Auth 레이아웃 nav 없이 생성 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteLayoutTSX_AuthLayout_NoNav(t *testing.T) {
	dir := t.TempDir()
	layout := stml.LayoutSpec{
		Name:      "auth",
		File:      "layouts/auth.html",
		HasOutlet: true,
	}
	if err := writeLayoutTSX(dir, layout, nil, "", nil); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(filepath.Join(dir, "layouts", "AuthLayout.tsx"))
	if err != nil {
		t.Fatal(err)
	}
	content := string(data)

	assertContains(t, content, "import { Outlet } from 'react-router-dom'")
	assertContains(t, content, "export default function AuthLayout()")
	assertContains(t, content, "<Outlet />")
	assertNotContains(t, content, "<nav>")
	assertNotContains(t, content, "Link")
}
