//ff:func feature=gen-react type=util control=iteration dimension=2
//ff:what resolveProtectedPages — 페이지가 소비하는 op 중 security 보호 op이 있으면 보호 페이지로 판정

package react

import (
	"github.com/park-jun-woo/yongol/pkg/validate/stml_openapi"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// resolveProtectedPages flags each STML page (by FileName) whose consumed
// ops — data-fetch/data-action blocks plus component api.<Op>( calls, the
// same per-page index XMO-10 trusts (stml_openapi.PageConsumedOps) — include
// at least one security-protected OpenAPI operation. Pages whose ops are all
// unsecured or security: [] stay public, so login/signup pages classify
// naturally without any hardcoded page-name list (Phase005). nil is returned
// when no OpenAPI document (no security truth source) or no backend.auth
// (no ProtectedRoute component to guard with) is present.
func resolveProtectedPages(fs *yongol.Fullstack) map[string]bool {
	if fs == nil || fs.OpenAPIDoc == nil {
		return nil
	}
	if hasAuth, _, _ := resolveAuthGates(fs); !hasAuth {
		return nil
	}
	secIndex := opSecurityIndex(fs.OpenAPIDoc)
	ops := make(map[string]struct{}, len(secIndex))
	for id := range secIndex {
		ops[id] = struct{}{}
	}
	out := make(map[string]bool, len(fs.STMLPages))
	for _, p := range fs.STMLPages {
		for id := range stml_openapi.PageConsumedOps(p, fs.SpecsDir, ops) {
			if secIndex[id] {
				out[p.FileName] = true
				break
			}
		}
	}
	return out
}
