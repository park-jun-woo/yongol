//ff:func feature=gen-hurl type=util control=sequence
//ff:what newScenarioCtx — 시나리오 빌드용 scenarioCtx 초기화 (manifest roles + shape-based auth op 감지)
package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// newScenarioCtx seeds captures with expected auth token names based on
// manifest roles, and runs SSaC shape-based auth op detection so every
// downstream phase shares the same signup/login decision.
func newScenarioCtx(fs *yongol.Fullstack) *scenarioCtx {
	ctx := &scenarioCtx{
		fs:         fs,
		captures:   map[string]bool{},
		roleMap:    buildOperationRoleMap(fs.ParsedPolicies),
		authOpIDs:  map[string]authRole{},
	}
	signup, login, _ := detectAuthOps(fs)
	ctx.authSignup = signup
	ctx.authLogin = login
	if signup != nil {
		ctx.authOpIDs[signup.OpID] = authRoleSignup
	}
	if login != nil {
		ctx.authOpIDs[login.OpID] = authRoleLogin
	}
	if fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return ctx
	}
	roles := fs.Manifest.Backend.Auth.Roles
	if len(roles) <= 1 {
		ctx.captures["token"] = true
		return ctx
	}
	seedRoleTokens(ctx, roles)
	return ctx
}
