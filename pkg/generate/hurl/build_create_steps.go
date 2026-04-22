//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what buildCreateSteps — Phase3a POST endpoint를 FK depth 순으로 step 생성 (state 전이 제외)
package hurl

// buildCreateSteps produces Create steps for POST endpoints sorted by FK depth.
// State-transition POSTs (operationIds that appear as stateDiagram events) are
// skipped here so that buildStateTransitions is the sole emitter with BFS order
// and branch-skip applied.
func buildCreateSteps(ctx *scenarioCtx) []step {
	fs := ctx.fs
	stateOps := collectStateOps(fs)
	sortedPaths := sortByFKDepth(fs.OpenAPIDoc.Paths, fs.DDLTables)
	var steps []step
	for _, sp := range sortedPaths {
		pathItem := fs.OpenAPIDoc.Paths.Find(sp.Path)
		if pathItem == nil || pathItem.Post == nil || pathItem.Post.OperationID == "" {
			continue
		}
		opID := pathItem.Post.OperationID
		if isAuthOpID(opID) || isPrereqResource(fs, resourceFromPath(sp.Path)) {
			continue
		}
		if stateOps[opID] {
			continue
		}
		s := buildPostStep(pathItem.Post, sp.Path, resourceFromPath(sp.Path), ctx)
		if len(steps) == 0 {
			s.Comment = "# ===== CRUD ====="
		}
		steps = append(steps, s)
	}
	return steps
}
