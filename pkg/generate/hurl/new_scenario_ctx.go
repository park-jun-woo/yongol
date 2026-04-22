//ff:func feature=gen-hurl type=util control=sequence
//ff:what newScenarioCtx — 시나리오 빌드용 scenarioCtx 초기화 (manifest roles → token captures)
package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// newScenarioCtx seeds captures with expected auth token names based on manifest roles.
func newScenarioCtx(fs *yongol.Fullstack) *scenarioCtx {
	ctx := &scenarioCtx{
		fs:       fs,
		captures: map[string]bool{},
		roleMap:  buildOperationRoleMap(fs.ParsedPolicies),
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
