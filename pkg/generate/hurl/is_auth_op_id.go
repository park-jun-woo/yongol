//ff:func feature=gen-hurl type=util control=sequence
//ff:what isAuthOpID — ctx.authOpIDs (SSaC shape 기반) 에 등록된 signup/login op 인지 판정
package hurl

// isAuthOpID returns true when opID was classified as signup or login
// by detectAuthOps during newScenarioCtx. Callers (create/read/update/
// delete step builders, auth FK collectors) use this to exclude auth
// endpoints from CRUD emission.
//
// The previous name-based allowlist ({"Register","register","Login",
// "login"}) was replaced in Phase003 (BUG-023) — classification is now
// driven by SSaC body shape (@verify-password, @call auth.HashPassword),
// so any operationId (Signup, Join, SignIn, EnrollStudent …) is handled
// uniformly.
func isAuthOpID(ctx *scenarioCtx, opID string) bool {
	if ctx == nil || ctx.authOpIDs == nil {
		return false
	}
	_, ok := ctx.authOpIDs[opID]
	return ok
}
