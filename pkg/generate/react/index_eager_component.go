//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what 인덱스 redirect 대상(indexTarget)이 가리키는 페이지의 컴포넌트명을 찾는다 (eager import 대상)

package react

// indexEagerComponent returns the component name of the route whose Path
// matches indexTarget — the index (first-entry) page that must stay eagerly
// imported so route-level lazy splitting (BUG-133) never delays the first
// paint. Returns "" when indexTarget is empty or no route matches, in which
// case every page is lazy-loaded.
func indexEagerComponent(routes []stmlRoute, indexTarget string) string {
	if indexTarget == "" {
		return ""
	}
	for _, r := range routes {
		if r.Path == indexTarget {
			return r.ComponentName
		}
	}
	return ""
}
