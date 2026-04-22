//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what seedRoleTokens — multi-role 모드에서 roles 목록을 "token_<role>" capture 키로 기록
package hurl

// seedRoleTokens records one "token_<role>" capture entry per role.
// Caller must ensure len(roles) >= 2 — single-role mode uses the plain
// "token" capture and should not reach here.
func seedRoleTokens(ctx *scenarioCtx, roles []string) {
	for _, role := range roles {
		ctx.captures["token_"+role] = true
	}
}
