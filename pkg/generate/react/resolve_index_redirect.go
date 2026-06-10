//ff:func feature=gen-react type=util control=iteration dimension=2
//ff:what resolveIndexRedirect — "/" 인덱스 redirect 대상(첫 공개 페이지, 없으면 /login) 결정

package react

import (
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// resolveIndexRedirect picks the Navigate target for the "/" index route
// (Phase005, BUG-111 (5)): the first public page in STML file-name sort
// order — the convention that decides "first page" — falling back to /login
// when every page is protected. Pages whose primary route carries a path
// param are skipped (a parameterized path is not a valid redirect target).
// When some page already routes "/" (data-route), that page is the index and
// "" is returned so no redirect route is emitted.
func resolveIndexRedirect(pages []stml.PageSpec, protectedPages map[string]bool) string {
	sorted := make([]stml.PageSpec, len(pages))
	copy(sorted, pages)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].FileName < sorted[j].FileName })

	for _, p := range sorted {
		for _, r := range pageToRoutes(p) {
			if r.Path == "/" {
				return ""
			}
		}
	}
	for _, p := range sorted {
		if protectedPages[p.FileName] {
			continue
		}
		rs := pageToRoutes(p)
		if len(rs) == 0 || strings.Contains(rs[0].Path, ":") {
			continue
		}
		return rs[0].Path
	}
	return "/login"
}
