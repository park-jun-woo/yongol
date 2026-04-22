//ff:func feature=gen-hurl type=util control=sequence
//ff:what buildPrereqSteps — Phase2 Prerequisite Creates: Auth body FK 의존 리소스 선생성
package hurl

// buildPrereqSteps creates prerequisite resources referenced by auth request bodies.
func buildPrereqSteps(ctx *scenarioCtx) []step {
	fs := ctx.fs
	if fs.Manifest == nil || fs.Manifest.Backend.Auth == nil || fs.OpenAPIDoc == nil {
		return nil
	}
	needed := collectAuthFKResources(fs)
	if len(needed) == 0 {
		return nil
	}
	return buildPrereqCreateSteps(ctx, needed)
}
