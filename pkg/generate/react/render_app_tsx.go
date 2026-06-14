//ff:func feature=gen-react type=generator control=iteration dimension=1
//ff:what stmlRoute 목록에서 전체 App.tsx 소스를 생성한다 (페이지별 가드 + "/" 인덱스 + catch-all)

package react

import (
	"fmt"
	"strings"
)

// renderAppTSX generates the full App.tsx source from a list of routes.
// layoutSet contains layout names that have LayoutSpec definitions.
// Routes with a non-empty Layout field are grouped under a layout wrapper
// route. Protected routes (page consumes a security-protected op) are
// wrapped with <ProtectedRoute> individually (Phase005 — replaces the
// blanket hasAuth wrap). indexTarget, when non-empty, emits a "/" index
// route redirecting there; a trailing catch-all path="*" redirects to "/"
// unless a page already claims "*".
func renderAppTSX(routes []stmlRoute, layoutSet map[string]bool, indexTarget string) string {
	grouped := groupRoutesByLayout(routes)

	anyProtected := false
	catchAll := true
	for _, r := range routes {
		if r.Protected {
			anyProtected = true
		}
		if r.Path == "*" {
			catchAll = false
		}
	}

	// Route-level code splitting (BUG-133): every page except the index page
	// is lazy-loaded. anyLazy gates the react lazy/Suspense import and the
	// <Suspense> wrapper so a project where every page is eager stays clean.
	eager := indexEagerComponent(routes, indexTarget)
	anyLazy := false
	for _, r := range routes {
		if r.ComponentName != eager {
			anyLazy = true
			break
		}
	}

	var sb strings.Builder
	if indexTarget != "" || catchAll {
		sb.WriteString("import { Routes, Route, Navigate } from 'react-router-dom'\n")
	} else {
		sb.WriteString("import { Routes, Route } from 'react-router-dom'\n")
	}
	if anyLazy {
		sb.WriteString("import { lazy, Suspense } from 'react'\n")
	}

	layoutNames := sortedLayoutNames(grouped)
	writeLayoutImports(&sb, layoutNames)

	if anyProtected {
		sb.WriteString("import ProtectedRoute from './components/ProtectedRoute'\n")
	}

	// Page imports last: eager static imports then lazy const declarations,
	// keeping every `import` statement ahead of any value binding.
	writePageImports(&sb, routes, indexTarget)

	sb.WriteString("\nexport default function App() {\n  return (\n")
	if anyLazy {
		sb.WriteString("    <Suspense fallback={<div>로딩 중...</div>}>\n")
	}
	sb.WriteString("    <Routes>\n")

	if indexTarget != "" {
		fmt.Fprintf(&sb, "      <Route path=\"/\" element={<Navigate to=\"%s\" replace />} />\n", indexTarget)
	}

	for _, name := range layoutNames {
		rs := grouped[name]
		if name == "" {
			writeFlatRoutes(&sb, rs)
			continue
		}
		writeLayoutGroupRoutes(&sb, name, rs)
	}

	if catchAll {
		sb.WriteString("      <Route path=\"*\" element={<Navigate to=\"/\" replace />} />\n")
	}

	sb.WriteString("    </Routes>\n")
	if anyLazy {
		sb.WriteString("    </Suspense>\n")
	}
	sb.WriteString("  )\n}\n")

	return sb.String()
}
