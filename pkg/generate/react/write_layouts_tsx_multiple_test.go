//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what writeLayoutsTSX 다중 레이아웃 파일 생성 검증

package react

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestWriteLayoutsTSX_Multiple(t *testing.T) {
	dir := t.TempDir()
	layouts := []stml.LayoutSpec{
		{
			Name: "app",
			NavItems: []stml.NavItem{
				{Path: "/workflows", Label: "Workflows"},
			},
			HasOutlet: true,
		},
		{
			Name:      "auth",
			HasOutlet: true,
		},
	}
	if err := writeLayoutsTSX(dir, layouts, nil, "", nil, "", "", nil); err != nil {
		t.Fatal(err)
	}

	for _, name := range []string{"AppLayout.tsx", "AuthLayout.tsx"} {
		if _, err := os.Stat(filepath.Join(dir, "layouts", name)); err != nil {
			t.Errorf("expected %s to exist: %v", name, err)
		}
	}
}
