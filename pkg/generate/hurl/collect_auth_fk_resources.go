//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what collectAuthFKResources — auth endpoint body에서 _id 접미사 FK 리소스 수집

package hurl

// collectAuthFKResources collects resource names from _id fields in auth request bodies.
// The caller passes the scenarioCtx whose authOpIDs set classifies which
// operations count as auth (shape-detected — Phase003).
func collectAuthFKResources(ctx *scenarioCtx) map[string]bool {
	needed := map[string]bool{}
	if ctx == nil || ctx.fs == nil || ctx.fs.OpenAPIDoc == nil {
		return needed
	}
	for _, pathItem := range ctx.fs.OpenAPIDoc.Paths.Map() {
		addAuthFKFields(ctx, pathItem, needed)
	}
	return needed
}
