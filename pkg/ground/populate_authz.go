//ff:func feature=rule type=loader control=sequence
//ff:what populateAuthz — built-in authz.CheckRequest 필드 집합을 Ground에 등록
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/rule"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// populateAuthz registers the default authz.CheckRequest fields into
// Ground.Lookup["Authz.checkRequest"]. XAS-60 (SSaC @auth input field ->
// CheckRequest) consumes this. When manifest declares a custom authz package,
// the set is left empty and XAS-60 is expected to skip.
func populateAuthz(g *rule.Ground, fs *yongol.Fullstack) {
	if fs.Manifest != nil && fs.Manifest.Authz != nil && fs.Manifest.Authz.Package != "" {
		g.Lookup["Authz.checkRequest"] = rule.StringSet{}
		return
	}
	g.Lookup["Authz.checkRequest"] = rule.StringSet{
		"Action":     true,
		"Resource":   true,
		"UserID":     true,
		"Role":       true,
		"ResourceID": true,
	}
}
