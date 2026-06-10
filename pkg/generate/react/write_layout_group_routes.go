//ff:func feature=gen-react type=emitter control=iteration dimension=1
//ff:what 레이아웃 그룹 라우트 JSX를 Builder에 기록한다 (보호 라우트는 페이지별 가드)

package react

import (
	"fmt"
	"strings"
)

// writeLayoutGroupRoutes writes a layout wrapper route with child routes.
// The wrapper itself is never guarded (Phase005 replaced the blanket
// layout-level wrap): each child route whose page consumes a
// security-protected op (r.Protected) gets its own <ProtectedRoute>, so
// public and protected pages can share one layout.
func writeLayoutGroupRoutes(sb *strings.Builder, name string, rs []stmlRoute) {
	compName := layoutComponentName(name)
	fmt.Fprintf(sb, "      <Route element={<%s />}>\n", compName)
	for _, r := range rs {
		if r.Protected {
			fmt.Fprintf(sb, "        <Route path=\"%s\" element={<ProtectedRoute><%s /></ProtectedRoute>} />\n", r.Path, r.ComponentName)
		} else {
			fmt.Fprintf(sb, "        <Route path=\"%s\" element={<%s />} />\n", r.Path, r.ComponentName)
		}
	}
	sb.WriteString("      </Route>\n")
}
