//ff:func feature=gen-react type=generator control=selection
//ff:what writeAuthStores — store 방출 게이트: claims 캡처 → claims store, bearer → 현행 store, 그 외 미방출 (hasRefresh로 refresh 필드 게이트)

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// writeAuthStores decides which src/stores/auth.ts shape the project gets
// (plans/stml/sitemap Phase005 — the store emission condition widened from
// "bearer" to "bearer ∨ claims captures exist"):
//   - an auth.claims.* capture exists → writeSessionStoreClaims, full
//     shape in bearer mode, claims-only in cookie mode; the store kind
//     falls back to "localStorage" when no frontend.auth resolved one
//     (claims captures without backend.auth);
//   - bearer without claims captures → writeSessionStore, byte-identical
//     to the pre-Phase005 output;
//   - otherwise (cookie/no-auth, no claims) → no store at all, as before.
func writeAuthStores(srcDir, authStore string, bearerAuth, hasRefresh bool, pages []stml.PageSpec) error {
	switch {
	case hasClaimsCaptures(pages):
		if authStore == "" {
			authStore = "localStorage"
		}
		return writeSessionStoreClaims(srcDir, authStore, bearerAuth, hasRefresh)
	case bearerAuth:
		return writeSessionStore(srcDir, authStore, hasRefresh)
	default:
		return nil
	}
}
