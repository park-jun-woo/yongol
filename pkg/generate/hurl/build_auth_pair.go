//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what buildAuthPair — Register→Login 순서로 step 생성 (BUG-015: 빈 DB에서도 실행 가능)

package hurl

// buildAuthPair builds auth steps in the order Register → Login for a
// given role. The explicit order replaces the previous map-iteration
// approach whose output depended on Go map randomization — which left
// Login (creds check) ahead of Register (creates the user) on typical
// zenflow runs, yielding 401 on an empty DB.
//
// Register comes first because:
//   - Empty DB → Register 201 creates a fresh account, Login then verifies.
//   - Both responses capture `access_token` into the same variable; the
//     Login-issued token overrides Register's, so downstream CRUD steps
//     always run with the freshest credential.
func buildAuthPair(ctx *scenarioCtx, role, tokenVar, sectionComment string) []step {
	if ctx.fs.OpenAPIDoc == nil {
		return nil
	}
	var out []step
	for _, opName := range authOpOrder() {
		s := findAuthOp(ctx, opName, role, tokenVar, sectionComment, len(out) == 0)
		if s != nil {
			out = append(out, *s)
		}
	}
	return out
}
