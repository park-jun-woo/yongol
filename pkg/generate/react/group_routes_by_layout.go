//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what 라우트를 레이아웃별로 그룹핑한다

package react

// groupRoutesByLayout partitions routes by their Layout field.
// Routes with empty Layout are keyed under "".
func groupRoutesByLayout(routes []stmlRoute) map[string][]stmlRoute {
	m := make(map[string][]stmlRoute)
	for _, r := range routes {
		m[r.Layout] = append(m[r.Layout], r)
	}
	return m
}
