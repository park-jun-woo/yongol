//ff:func feature=gen-react type=emitter control=iteration dimension=1
//ff:what 페이지 컴포넌트 import 문을 Builder에 기록한다 (인덱스 페이지만 eager, 나머지는 lazy)

package react

import (
	"fmt"
	"strings"
)

// writePageImports writes page component imports, deduplicating by component
// name. The index page (the route matching indexTarget) is emitted as a
// static eager import so it never blocks the first paint; every other page is
// emitted as a `const X = lazy(() => import(...))` declaration for route-level
// code splitting (BUG-133). Eager static imports are written before the lazy
// const declarations so all `import` statements precede any value binding.
func writePageImports(sb *strings.Builder, routes []stmlRoute, indexTarget string) {
	eager := indexEagerComponent(routes, indexTarget)
	seen := make(map[string]bool)

	for _, r := range routes {
		if seen[r.ComponentName] || r.ComponentName != eager {
			continue
		}
		seen[r.ComponentName] = true
		fmt.Fprintf(sb, "import %s from '%s'\n", r.ComponentName, r.ImportPath)
	}
	for _, r := range routes {
		if seen[r.ComponentName] {
			continue
		}
		seen[r.ComponentName] = true
		fmt.Fprintf(sb, "const %s = lazy(() => import('%s'))\n", r.ComponentName, r.ImportPath)
	}
}
