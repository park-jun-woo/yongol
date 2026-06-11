//ff:func feature=gen-react type=util control=sequence
//ff:what sortBreadcrumbRoutes — 라우트 매칭 테이블 정렬 (파라미터 세그먼트 적은 패턴 우선, 동률은 문서 순서 유지)

package react

import (
	"sort"
	"strings"
)

// sortBreadcrumbRoutes orders the BREADCRUMB_ROUTES matching table so the
// generated component's first-match scan approximates react-router's
// route ranking: patterns with fewer parameter segments come first
// (a static "/buildings/new" must win over "/buildings/:BuildingID" for
// the pathname "/buildings/new"), ties keep sitemap document order
// (stable sort). The input slice is not mutated.
func sortBreadcrumbRoutes(trails []breadcrumbTrail) []breadcrumbTrail {
	sorted := append([]breadcrumbTrail{}, trails...)
	sort.SliceStable(sorted, func(i, j int) bool {
		return strings.Count(sorted[i].Pattern, ":") < strings.Count(sorted[j].Pattern, ":")
	})
	return sorted
}
