//ff:func feature=gen-react type=generator control=sequence
//ff:what App.tsx — STML 페이지 목록에서 React Router 라우트 자동 생성 (페이지별 보호 가드 포함)

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
// protectedPages (Phase005 — resolveProtectedPages, keyed by FileName) flags
// pages whose ops carry OpenAPI security; their routes are wrapped with
// <ProtectedRoute> per page. A "/" index route redirecting to the declared
// index page (indexPage — manifest frontend.index, page-flow Phase009) or,
// when undeclared, to the first public page, and a catch-all path="*" are
// emitted alongside (BUG-111 (5)).
func writeAppTSX(srcDir string, pages []stml.PageSpec, layouts []stml.LayoutSpec, defaultLayout string, protectedPages map[string]bool, indexPage string) error {
	if len(pages) == 0 {
		return writeAppTSXPlaceholder(srcDir)
	}

	routes := buildRoutes(pages, defaultLayout, protectedPages)
	layoutSet := buildLayoutSet(layouts)
	indexTarget := resolveIndexRedirect(pages, protectedPages, indexPage)
	src := renderAppTSX(routes, layoutSet, indexTarget)
	return os.WriteFile(filepath.Join(srcDir, "App.tsx"), []byte(src), 0644)
}
