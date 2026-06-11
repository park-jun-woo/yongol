//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what staticRoutePrefix — 라우트 패턴을 첫 파라미터 세그먼트 앞에서 잘라 정적 prefix 반환 (조상 하이라이트 매칭용)

package react

import "strings"

// staticRoutePrefix cuts a route pattern at its first parameter segment,
// returning the static prefix the ancestor-highlight matcher tests with
// pathname.startsWith (plans/stml/sitemap Phase003, DESIGN §4.4). A
// parameterized pattern keeps a trailing "/" so "/buildings/:ID" matches
// "/buildings/7" but not the list route "/buildings" itself (exact
// matching is NavLink end's job); a fully static pattern is returned
// verbatim. "" in, "" out; a leading parameter yields "/" — both are
// useless prefixes the caller must skip.
func staticRoutePrefix(pattern string) string {
	segs := strings.Split(pattern, "/")
	for i, s := range segs {
		if strings.HasPrefix(s, ":") {
			return strings.Join(segs[:i], "/") + "/"
		}
	}
	return pattern
}
