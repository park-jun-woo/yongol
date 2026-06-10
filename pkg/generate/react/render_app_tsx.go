//ff:func feature=gen-react type=generator control=iteration dimension=1
//ff:what stmlRoute 목록에서 전체 App.tsx 소스를 생성한다

package react

import (
	"strings"
)

// renderAppTSX generates the full App.tsx source from a list of routes.
// layoutSet contains layout names that have LayoutSpec definitions.
// Routes with a non-empty Layout field are grouped under a layout wrapper route.
// When hasAuth is true, non-auth layout groups and flat routes are wrapped
// with <ProtectedRoute>.
func renderAppTSX(routes []stmlRoute, layoutSet map[string]bool, hasAuth bool) string {
	grouped := groupRoutesByLayout(routes)

	var sb strings.Builder
	sb.WriteString("import { Routes, Route } from 'react-router-dom'\n")

	writePageImports(&sb, routes)

	layoutNames := sortedLayoutNames(grouped)
	writeLayoutImports(&sb, layoutNames)

	if hasAuth {
		sb.WriteString("import ProtectedRoute from './components/ProtectedRoute'\n")
	}

	sb.WriteString("\nexport default function App() {\n  return (\n    <Routes>\n")

	for _, name := range layoutNames {
		rs := grouped[name]
		if name == "" {
			writeFlatRoutes(&sb, rs, hasAuth)
			continue
		}
		writeLayoutGroupRoutes(&sb, name, rs, hasAuth)
	}

	sb.WriteString("    </Routes>\n  )\n}\n")

	return sb.String()
}
