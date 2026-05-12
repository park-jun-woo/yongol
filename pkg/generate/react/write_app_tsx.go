//ff:func feature=gen-react type=generator control=sequence
//ff:what App.tsx — STML 페이지 목록에서 React Router 라우트 자동 생성

package react

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// stmlRoute represents a single route derived from an STML page.
type stmlRoute struct {
	Path          string // e.g. "/workflows/:id"
	ComponentName string // e.g. "WorkflowDetail"
	ImportPath    string // e.g. "./pages/workflow-detail"
	Layout        string // layout name (e.g. "app", "auth"); empty = no layout
}

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

// writeAppTSXPlaceholder emits a minimal placeholder App.tsx when no STML
// pages are available.
func writeAppTSXPlaceholder(srcDir string) error {
	const src = `import { Routes, Route } from 'react-router-dom'

// Add your pages under src/pages/*.tsx and wire them below.
//   e.g. <Route path="/workflows" element={<WorkflowsPage />} />
export default function App() {
  return (
    <Routes>
      <Route path="/" element={<div className="p-6">yongol scaffolded frontend — add pages under src/pages/</div>} />
    </Routes>
  )
}
`
	return os.WriteFile(filepath.Join(srcDir, "App.tsx"), []byte(src), 0644)
}

// buildRoutes converts STML PageSpecs into sorted route definitions.
// defaultLayout is applied to pages that have no explicit Layout set.
func buildRoutes(pages []stml.PageSpec, defaultLayout string) []stmlRoute {
	routes := make([]stmlRoute, 0, len(pages))
	for _, p := range pages {
		rs := pageToRoutes(p)
		// Resolve layout: explicit page.Layout > defaultLayout > ""
		resolvedLayout := p.Layout
		if resolvedLayout == "" {
			resolvedLayout = defaultLayout
		}
		for i := range rs {
			rs[i].Layout = resolvedLayout
		}
		routes = append(routes, rs...)
	}
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Path < routes[j].Path
	})
	return routes
}

// buildLayoutSet returns a set of layout names that have corresponding LayoutSpecs.
func buildLayoutSet(layouts []stml.LayoutSpec) map[string]bool {
	s := make(map[string]bool, len(layouts))
	for _, l := range layouts {
		s[l.Name] = true
	}
	return s
}

// pageToRoutes converts a single STML PageSpec into route definitions.
//
// Rules:
//  0. If page.Route is set (data-route), use it as-is (single route)
//  1. Strip .html → kebab-case path (e.g. "workflows.html" → "/workflows")
//  2. "-detail" suffix → parent resource path + /:id (single route)
//     e.g. "workflow-detail.html" → "/workflows/:id"
//  3. Non-detail page with route param → two routes: base path + base/:id
//     e.g. "templates.html" with data-param-id="route.id" → "/templates" and "/templates/:id"
func pageToRoutes(p stml.PageSpec) []stmlRoute {
	base := strings.TrimSuffix(p.FileName, ".html")
	componentName := kebabToPascal(base)
	importPath := "./pages/" + base

	// Explicit data-route takes priority over filename-based inference.
	if p.Route != "" {
		return []stmlRoute{{
			Path:          p.Route,
			ComponentName: componentName,
			ImportPath:    importPath,
		}}
	}

	hasRouteParam := pageHasRouteParam(p)

	if strings.HasSuffix(base, "-detail") {
		// e.g. "workflow-detail" → parent "workflow" → pluralize → "/workflows/:id"
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
		// Non-detail page with route param: emit both base and parameterized route
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

// pageHasRouteParam returns true if any fetch or action block in the page
// contains a data-param-* attribute with a "route." prefixed source.
func pageHasRouteParam(p stml.PageSpec) bool {
	for _, f := range p.Fetches {
		for _, param := range f.Params {
			if strings.HasPrefix(param.Source, "route.") {
				return true
			}
		}
	}
	for _, a := range p.Actions {
		for _, param := range a.Params {
			if strings.HasPrefix(param.Source, "route.") {
				return true
			}
		}
	}
	return false
}

// kebabToPascal converts a kebab-case string to PascalCase.
// e.g. "workflow-detail" → "WorkflowDetail"
func kebabToPascal(s string) string {
	parts := strings.Split(s, "-")
	for i, part := range parts {
		if len(part) > 0 {
			parts[i] = strings.ToUpper(part[:1]) + part[1:]
		}
	}
	return strings.Join(parts, "")
}

// naivePluralize adds "s" to a kebab-case string for simple English
// pluralization. Handles common suffixes: -s, -sh, -ch, -x, -z → "es".
func naivePluralize(s string) string {
	if s == "" {
		return s
	}
	if strings.HasSuffix(s, "s") || strings.HasSuffix(s, "sh") ||
		strings.HasSuffix(s, "ch") || strings.HasSuffix(s, "x") ||
		strings.HasSuffix(s, "z") {
		return s + "es"
	}
	return s + "s"
}

// isAuthLayout returns true if the layout name is "auth".
// Auth layouts host public pages (login, register) and are never wrapped
// with ProtectedRoute.
func isAuthLayout(name string) bool {
	return name == "auth"
}

// renderAppTSX generates the full App.tsx source from a list of routes.
// layoutSet contains layout names that have LayoutSpec definitions.
// Routes with a non-empty Layout field are grouped under a layout wrapper route.
// When hasAuthz is true, non-auth layout groups and flat routes are wrapped
// with <ProtectedRoute>.
func renderAppTSX(routes []stmlRoute, layoutSet map[string]bool, hasAuthz bool) string {
	// Partition routes by layout
	grouped := groupRoutesByLayout(routes)

	var sb strings.Builder
	sb.WriteString("import { Routes, Route } from 'react-router-dom'\n")

	// Import page components
	seen := make(map[string]bool)
	for _, r := range routes {
		if seen[r.ComponentName] {
			continue
		}
		seen[r.ComponentName] = true
		fmt.Fprintf(&sb, "import %s from '%s'\n", r.ComponentName, r.ImportPath)
	}

	// Import layout components (sorted for deterministic output)
	layoutNames := sortedLayoutNames(grouped)
	for _, name := range layoutNames {
		if name == "" {
			continue
		}
		compName := layoutComponentName(name)
		fmt.Fprintf(&sb, "import %s from './layouts/%s'\n", compName, compName)
	}

	// Import ProtectedRoute when authz is enabled
	if hasAuthz {
		sb.WriteString("import ProtectedRoute from './components/ProtectedRoute'\n")
	}

	sb.WriteString("\nexport default function App() {\n  return (\n    <Routes>\n")

	// Emit layout-grouped routes first, then flat routes
	for _, name := range layoutNames {
		rs := grouped[name]
		if name == "" {
			// Flat routes (no layout)
			if hasAuthz {
				// Wrap each flat route with ProtectedRoute
				for _, r := range rs {
					fmt.Fprintf(&sb, "      <Route path=\"%s\" element={<ProtectedRoute><%s /></ProtectedRoute>} />\n", r.Path, r.ComponentName)
				}
			} else {
				for _, r := range rs {
					fmt.Fprintf(&sb, "      <Route path=\"%s\" element={<%s />} />\n", r.Path, r.ComponentName)
				}
			}
			continue
		}
		compName := layoutComponentName(name)
		if hasAuthz && !isAuthLayout(name) {
			// Wrap the layout element with ProtectedRoute
			fmt.Fprintf(&sb, "      <Route element={<ProtectedRoute><%s /></ProtectedRoute>}>\n", compName)
		} else {
			fmt.Fprintf(&sb, "      <Route element={<%s />}>\n", compName)
		}
		for _, r := range rs {
			fmt.Fprintf(&sb, "        <Route path=\"%s\" element={<%s />} />\n", r.Path, r.ComponentName)
		}
		sb.WriteString("      </Route>\n")
	}

	sb.WriteString("    </Routes>\n  )\n}\n")

	return sb.String()
}

// groupRoutesByLayout partitions routes by their Layout field.
// Routes with empty Layout are keyed under "".
func groupRoutesByLayout(routes []stmlRoute) map[string][]stmlRoute {
	m := make(map[string][]stmlRoute)
	for _, r := range routes {
		m[r.Layout] = append(m[r.Layout], r)
	}
	return m
}

// sortedLayoutNames returns layout names in sorted order, with "" (flat) last.
func sortedLayoutNames(grouped map[string][]stmlRoute) []string {
	names := make([]string, 0, len(grouped))
	hasFlat := false
	for name := range grouped {
		if name == "" {
			hasFlat = true
			continue
		}
		names = append(names, name)
	}
	sort.Strings(names)
	if hasFlat {
		names = append(names, "")
	}
	return names
}
