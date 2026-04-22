//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what buildAuthSteps — Phase1 Auth: Register+Login 고정 순서, JWT token capture
package hurl

// buildAuthSteps produces Register + Login steps if manifest.auth exists.
// Token captures (token, token_{role}) are already seeded into ctx.captures
// by newScenarioCtx so downstream phases can resolve tokens.
func buildAuthSteps(ctx *scenarioCtx) []step {
	fs := ctx.fs
	if fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return nil
	}
	roles := fs.Manifest.Backend.Auth.Roles
	sectionComment := "# ===== Auth ====="

	if len(roles) <= 1 {
		return buildAuthPair(ctx, "", "token", sectionComment)
	}

	var steps []step
	for i, role := range roles {
		comment := ""
		if i == 0 {
			comment = sectionComment
		}
		steps = append(steps, buildAuthPair(ctx, role, "token_"+role, comment)...)
	}
	return steps
}
