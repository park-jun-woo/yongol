//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what indexFallbackPage — 인덱스 미선언 시 에미터 폴백(파일명 정렬 첫 공개 페이지)을 동일 규칙으로 재현

package stml_openapi

import (
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// indexFallbackPage mirrors the emitter's index fallback
// (react.resolveIndexRedirect step 3) so TM-35 can name the page the
// fallback actually picks: the first page in STML file-name sort order
// that is public (consumes no security-protected operation — the same
// judgment react.resolveProtectedPages applies, gated on backend.auth
// presence) and whose primary resolved route carries no path param.
// Returns the page file name and its route; ("", "/login") when every
// candidate is protected or parameterized.
func indexFallbackPage(fs *yongol.Fullstack, opMap map[string]operationEntry) (string, string) {
	sorted := make([]stml.PageSpec, len(fs.STMLPages))
	copy(sorted, fs.STMLPages)
	sort.Slice(sorted, func(i, j int) bool { return sorted[i].FileName < sorted[j].FileName })

	authActive := backendAuthMode(fs) != ""
	ops := make(map[string]struct{}, len(opMap))
	for id := range opMap {
		ops[id] = struct{}{}
	}

	for _, p := range sorted {
		if authActive && pageConsumesProtectedOp(p, fs, opMap, ops) {
			continue
		}
		patterns := stml.RoutePaths(p)
		if len(patterns) == 0 || strings.Contains(patterns[0], ":") {
			continue
		}
		return p.FileName, patterns[0]
	}
	return "", "/login"
}
