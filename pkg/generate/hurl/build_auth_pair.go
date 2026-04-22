//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what buildAuthPair — Register+Login 쌍의 step 생성 (역할별 token capture)
package hurl

// buildAuthPair builds Register + Login steps for a given role.
func buildAuthPair(ctx *scenarioCtx, role, tokenVar, sectionComment string) []step {
	if ctx.fs.OpenAPIDoc == nil {
		return nil
	}
	var out []step
	for path, pathItem := range ctx.fs.OpenAPIDoc.Paths.Map() {
		s := matchAuthOp(pathItem, path, ctx, role, tokenVar, sectionComment, len(out) == 0)
		if s != nil {
			out = append(out, *s)
		}
	}
	return out
}
