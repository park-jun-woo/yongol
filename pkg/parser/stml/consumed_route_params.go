//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what ConsumedRouteParams — 페이지가 소비하는 route.<Name> 이름을 첫 등장 순서로 반환 (TM-27/28 소비)
package stml

// ConsumedRouteParams returns the names of the route.<Name> params the
// page consumes, in first-appearance order. It is the exported view of
// collectRouteParams for cross-package consumers: TM-27/28 compare these
// names against the resolved route patterns of RoutePaths.
func ConsumedRouteParams(p PageSpec) []string {
	params := collectRouteParams(p)
	if len(params) == 0 {
		return nil
	}
	names := make([]string, 0, len(params))
	for _, rp := range params {
		names = append(names, rp.Name)
	}
	return names
}
