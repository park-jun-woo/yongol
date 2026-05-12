//ff:func feature=gen-react type=emitter control=iteration dimension=1
//ff:what 페이지 컴포넌트 import 문을 Builder에 기록한다

package react

import (
	"fmt"
	"strings"
)

// writePageImports writes page component import statements, deduplicating by component name.
func writePageImports(sb *strings.Builder, routes []stmlRoute) {
	seen := make(map[string]bool)
	for _, r := range routes {
		if seen[r.ComponentName] {
			continue
		}
		seen[r.ComponentName] = true
		fmt.Fprintf(sb, "import %s from '%s'\n", r.ComponentName, r.ImportPath)
	}
}
