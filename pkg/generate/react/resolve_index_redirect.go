//ff:func feature=gen-react type=util control=iteration dimension=2
//ff:what resolveIndexRedirect — "/" 인덱스 redirect 대상 결정 (data-route="/" > frontend.index 선언 > 첫 공개 페이지 폴백)

package react

import (
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// resolveIndexRedirect picks the Navigate target for the "/" index route.
// Priority (page-flow Phase009, BUG-114 (3)):
//
//  1. A page that routes "/" itself (data-route) — that page is the index,
//     "" is returned so no redirect route is emitted (Phase005 behaviour).
//  2. indexPage (manifest frontend.index, an STML page name) — the target
//     page's resolved route with optional segments (":Name?") stripped:
//     a redirect has no value to fill them, and TM-34 permits optional-only
//     pages, so the literal must not leak into <Navigate to>. Required
//     segments cannot appear here — TM-34 blocks them before generate.
//     A protected index page is legal: <ProtectedRoute> bounces an
//     unauthenticated visit to /login after the redirect.
//  3. Fallback: the first public page in STML file-name sort order — the
//     convention that decides "first page" — falling back to /login when
//     every page is protected. Pages whose primary route carries a path
//     param are skipped (a parameterized path is not a valid redirect
//     target). TM-35 surfaces this fallback as a WARNING.
func resolveIndexRedirect(pages []stml.PageSpec, protectedPages map[string]bool, indexPage string) string {
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
	if indexPage != "" {
		for _, p := range sorted {
			if p.Name != indexPage {
				continue
			}
			if rs := pageToRoutes(p); len(rs) > 0 {
				return stripOptionalSegments(rs[0].Path)
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
