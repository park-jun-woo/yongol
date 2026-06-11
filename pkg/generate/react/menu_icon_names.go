//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what menuIconNames — 메뉴 항목 트리에서 사용된 lucide-react 컴포넌트명 수집 (정렬·중복 제거, import 라인용)

package react

import "sort"

// menuIconNames gathers the lucide-react component names referenced by the
// rendered menu items, deduplicated and sorted for a deterministic import
// line. Only rendered items reach the model, so the import set matches the
// emitted JSX exactly — tsc flags an unknown icon name as an import error
// (no silent runtime fallback).
func menuIconNames(items []sitemapMenuItem) []string {
	seen := map[string]struct{}{}
	stack := append([]sitemapMenuItem(nil), items...)
	for len(stack) > 0 {
		it := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if it.Icon != "" {
			seen[it.Icon] = struct{}{}
		}
		stack = append(stack, it.Children...)
	}
	if len(seen) == 0 {
		return nil
	}
	names := make([]string, 0, len(seen))
	for n := range seen {
		names = append(names, n)
	}
	sort.Strings(names)
	return names
}
