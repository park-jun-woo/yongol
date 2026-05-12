//ff:func feature=gen-react type=generator control=sequence
//ff:what App.tsx — STML 페이지 목록에서 React Router 라우트 자동 생성

package react

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// writeAppTSX emits App.tsx with routes derived from STML pages.
// If no STML pages are provided, a placeholder scaffold is emitted.
//
// When layouts are provided, routes are grouped under layout wrapper routes:
//   - page.Layout != "" → grouped under that layout
//   - page.Layout == "" → grouped under defaultLayout (if non-empty)
//   - defaultLayout == "" and page.Layout == "" → flat route (no wrapper)
//
// When hasAuthz is true, non-auth layout groups and flat routes are wrapped
// with <ProtectedRoute>. The "auth" layout is always public.
func writeAppTSX(srcDir string, pages []stml.PageSpec, layouts []stml.LayoutSpec, defaultLayout string, hasAuthz bool) error {
	if len(pages) == 0 {
		return writeAppTSXPlaceholder(srcDir)
	}

	routes := buildRoutes(pages, defaultLayout)
	layoutSet := buildLayoutSet(layouts)
	src := renderAppTSX(routes, layoutSet, hasAuthz)
	return os.WriteFile(filepath.Join(srcDir, "App.tsx"), []byte(src), 0644)
}
