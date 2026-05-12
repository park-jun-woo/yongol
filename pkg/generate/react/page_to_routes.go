//ff:func feature=gen-react type=util control=sequence
//ff:what 단일 STML PageSpec을 라우트 정의로 변환한다

package react

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// pageToRoutes converts a single STML PageSpec into route definitions.
//
// Rules:
//
//	0. If page.Route is set (data-route), use it as-is (single route)
//	1. Strip .html → kebab-case path (e.g. "workflows.html" → "/workflows")
//	2. "-detail" suffix → parent resource path + /:id (single route)
//	   e.g. "workflow-detail.html" → "/workflows/:id"
//	3. Non-detail page with route param → two routes: base path + base/:id
//	   e.g. "templates.html" with data-param-id="route.id" → "/templates" and "/templates/:id"
func pageToRoutes(p stml.PageSpec) []stmlRoute {
	base := strings.TrimSuffix(p.FileName, ".html")
	componentName := kebabToPascal(base)
	importPath := "./pages/" + base

	if p.Route != "" {
		return []stmlRoute{{
			Path:          p.Route,
			ComponentName: componentName,
			ImportPath:    importPath,
		}}
	}

	hasRouteParam := pageHasRouteParam(p)

	if strings.HasSuffix(base, "-detail") {
		parent := strings.TrimSuffix(base, "-detail")
		parentPath := "/" + naivePluralize(parent)
		return []stmlRoute{{
			Path:          parentPath + "/:id",
			ComponentName: componentName,
			ImportPath:    importPath,
		}}
	}

	routePath := "/" + base
	if hasRouteParam {
		return []stmlRoute{
			{Path: routePath, ComponentName: componentName, ImportPath: importPath},
			{Path: routePath + "/:id", ComponentName: componentName, ImportPath: importPath},
		}
	}

	return []stmlRoute{{
		Path:          routePath,
		ComponentName: componentName,
		ImportPath:    importPath,
	}}
}
