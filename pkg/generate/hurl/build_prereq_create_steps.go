//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what buildPrereqCreateSteps — needed 리소스에 해당하는 POST endpoint step 생성
package hurl

// buildPrereqCreateSteps creates steps for POST endpoints that produce needed resources.
func buildPrereqCreateSteps(ctx *scenarioCtx, needed map[string]bool) []step {
	var steps []step
	for path, pathItem := range ctx.fs.OpenAPIDoc.Paths.Map() {
		if pathItem.Post == nil || pathItem.Post.OperationID == "" {
			continue
		}
		resource := resourceFromPath(path)
		if !needed[resource] {
			continue
		}
		s := buildPostStep(pathItem.Post, path, resource, ctx)
		if len(steps) == 0 {
			s.Comment = "# ===== Prereq ====="
		}
		steps = append(steps, s)
	}
	return steps
}
