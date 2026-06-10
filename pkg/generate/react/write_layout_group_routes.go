//ff:func feature=gen-react type=emitter control=iteration dimension=1
//ff:what 레이아웃 그룹 라우트 JSX를 Builder에 기록한다

package react

import (
	"fmt"
	"strings"
)

// writeLayoutGroupRoutes writes a layout wrapper route with child routes.
func writeLayoutGroupRoutes(sb *strings.Builder, name string, rs []stmlRoute, hasAuth bool) {
	compName := layoutComponentName(name)
	if hasAuth && !isAuthLayout(name) {
		fmt.Fprintf(sb, "      <Route element={<ProtectedRoute><%s /></ProtectedRoute>}>\n", compName)
	} else {
		fmt.Fprintf(sb, "      <Route element={<%s />}>\n", compName)
	}
	for _, r := range rs {
		fmt.Fprintf(sb, "        <Route path=\"%s\" element={<%s />} />\n", r.Path, r.ComponentName)
	}
	sb.WriteString("      </Route>\n")
}
