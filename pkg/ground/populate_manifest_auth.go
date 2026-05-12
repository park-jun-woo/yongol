//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateManifestAuth — manifest auth 섹션에서 claims, roles 설정 추출하여 Ground 등록

package ground

import (
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func populateManifestAuth(g *rule.Ground, auth *manifest.Auth) {
	claims := make(rule.StringSet)
	claimKeys := make(rule.StringSet)
	for field, def := range auth.Claims {
		claims[field] = true
		claimKeys[def.Key] = true
		goType := def.GoType
		if goType == "uuid" {
			goType = "pgtype.UUID"
		}
		g.Types["Manifest.claim."+field] = goType
	}
	g.Lookup["Manifest.claims"] = claims
	g.Lookup["Manifest.claims.keys"] = claimKeys
	g.Config["backend.auth.claims"] = len(claims) > 0

	roles := make(rule.StringSet)
	for _, r := range auth.Roles {
		roles[r] = true
	}
	g.Lookup["Manifest.roles"] = roles
}
