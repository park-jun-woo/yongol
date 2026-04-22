//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what populateFunc — FuncSpec 를 Ground 의 Func.spec / Func.request / Struct.* 에 등록
package ground

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// populateFunc registers all FuncSpecs (project + yongol built-in) into Ground
// lookup stores. Detail logic for a single spec lives in registerFuncSpec.
// camelCase @func annotation name is the canonical key; SSaC @call resolution
// is the caller's responsibility.
func populateFunc(g *rule.Ground, fs *yongol.Fullstack) {
	specs := make(rule.StringSet)
	allSpecs := append(fs.ProjectFuncSpecs, fs.YongolPkgSpecs...)
	for i := range allSpecs {
		sp := &allSpecs[i]
		specs[sp.Package+"."+sp.Name] = true
		registerFuncSpec(g, sp)
	}
	if fs.Manifest != nil && fs.Manifest.Backend.Auth != nil && len(fs.Manifest.Backend.Auth.Claims) > 0 {
		specs["auth.issueToken"] = true
		specs["auth.verifyToken"] = true
		specs["auth.refreshToken"] = true
		// Phase009 — SSaC 정규 경로용 고수준 함수 (POST /auth/refresh, POST /auth/logout).
		specs["auth.refreshRotate"] = true
		specs["auth.logout"] = true
	}
	g.Lookup["Func.spec"] = specs
}
