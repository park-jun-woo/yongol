//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what 단일 STML PageSpec을 stml.RoutePaths가 유도한 경로로 라우트 정의로 변환한다

package react

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// pageToRoutes converts a single STML PageSpec into route definitions.
// Route planning lives in stml.RoutePaths — the single source for the
// explicit data-route override and the paths derived from the page's
// route.<Name> consumption — so the emitted route patterns and the page's
// useParams() destructuring share one table. This function only attaches
// the component name and import path.
func pageToRoutes(p stml.PageSpec) []stmlRoute {
	base := strings.TrimSuffix(p.FileName, ".html")
	componentName := kebabToPascal(base)
	importPath := "./pages/" + base

	paths := stml.RoutePaths(p)
	routes := make([]stmlRoute, 0, len(paths))
	for _, path := range paths {
		routes = append(routes, stmlRoute{
			Path:          path,
			ComponentName: componentName,
			ImportPath:    importPath,
		})
	}
	return routes
}
