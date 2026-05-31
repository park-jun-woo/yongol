//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateMiddleware — bearerAuth 사용 시 manifest claim KEY를 Lookup["Middleware.claims"]에 등록
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// populateMiddleware registers JWT claim KEY names that the auto-generated
// bearerAuth middleware consumes. When manifest.backend.middleware contains
// "bearerAuth" AND backend.auth.claims is configured, every claim KEY is
// added to Lookup["Middleware.claims"].
//
// XPN-54 consults this set to avoid false-positive "claim not referenced"
// warnings for claims that exist solely to feed the auto-generated middleware
// (e.g. `ID: user_id` is consumed by bearerAuth even when no Rego rule
// references input.claims.user_id directly).
func populateMiddleware(g *rule.Ground, fs *yongol.Fullstack) {
	if fs.Manifest == nil {
		return
	}
	if !hasBearerAuthMiddleware(fs.Manifest.Backend.Middleware) {
		return
	}
	if fs.Manifest.Backend.Auth == nil {
		return
	}
	set := make(rule.StringSet)
	for _, c := range fs.Manifest.Backend.Auth.Claims {
		if c.Key != "" {
			set[c.Key] = true
		}
	}
	g.Lookup["Middleware.claims"] = set
}
