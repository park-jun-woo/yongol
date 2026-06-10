//ff:func feature=gen-react type=emitter control=iteration dimension=1
//ff:what 레이아웃 없는 flat 라우트 JSX를 Builder에 기록한다

package react

import (
	"fmt"
	"strings"
)

// writeFlatRoutes writes flat (no layout) route elements.
func writeFlatRoutes(sb *strings.Builder, rs []stmlRoute, hasAuth bool) {
	for _, r := range rs {
		if hasAuth {
			fmt.Fprintf(sb, "      <Route path=\"%s\" element={<ProtectedRoute><%s /></ProtectedRoute>} />\n", r.Path, r.ComponentName)
		} else {
			fmt.Fprintf(sb, "      <Route path=\"%s\" element={<%s />} />\n", r.Path, r.ComponentName)
		}
	}
}
