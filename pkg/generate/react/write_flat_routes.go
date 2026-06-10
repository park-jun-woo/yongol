//ff:func feature=gen-react type=emitter control=iteration dimension=1
//ff:what 레이아웃 없는 flat 라우트 JSX를 Builder에 기록한다 (보호 라우트는 페이지별 가드)

package react

import (
	"fmt"
	"strings"
)

// writeFlatRoutes writes flat (no layout) route elements. Routes whose page
// consumes a security-protected op (r.Protected, Phase005) are wrapped with
// <ProtectedRoute> individually — public pages render unguarded.
func writeFlatRoutes(sb *strings.Builder, rs []stmlRoute) {
	for _, r := range rs {
		if r.Protected {
			fmt.Fprintf(sb, "      <Route path=\"%s\" element={<ProtectedRoute><%s /></ProtectedRoute>} />\n", r.Path, r.ComponentName)
		} else {
			fmt.Fprintf(sb, "      <Route path=\"%s\" element={<%s />} />\n", r.Path, r.ComponentName)
		}
	}
}
