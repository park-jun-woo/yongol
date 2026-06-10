//ff:func feature=gen-react type=generator control=sequence
//ff:what LayoutSpec -> React layout TSX 파일 생성 (Link + Outlet + 모드별 로그아웃)

package react

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// writeLayoutTSX emits a single layout TSX file from a LayoutSpec.
// The file is written to <srcDir>/layouts/<PascalName>Layout.tsx.
//
// Conversion rules:
//   - NavItem → <Link to="path">Label</Link> (page-name references
//     resolve through routePatterns — page-flow Phase010)
//   - HasOutlet → <Outlet />
//   - Logout → mode-wired logout button (authMode "bearer"/"cookie";
//     "" = no auth, emission skipped)
//   - Layout name → PascalCase + "Layout" suffix (e.g. "app" → "AppLayout")
func writeLayoutTSX(srcDir string, layout stml.LayoutSpec, routePatterns map[string]string, authMode string) error {
	layoutsDir := filepath.Join(srcDir, "layouts")
	if err := os.MkdirAll(layoutsDir, 0o755); err != nil {
		return err
	}

	componentName := layoutComponentName(layout.Name)
	src := renderLayoutTSX(componentName, layout, routePatterns, authMode)
	return os.WriteFile(filepath.Join(layoutsDir, componentName+".tsx"), []byte(src), 0o644)
}
