//ff:func feature=gen-hurl type=util control=sequence
//ff:what buildDeleteSteps — Phase5 Cleanup Deletes: FK 역순 (자식 먼저) DELETE endpoint
package hurl

// buildDeleteSteps produces Phase 5 cleanup steps: DELETE endpoints in FK reverse order.
func buildDeleteSteps(ctx *scenarioCtx) []step {
	fs := ctx.fs
	if fs.OpenAPIDoc == nil {
		return nil
	}
	deletes := collectDeleteOps(ctx)
	sortDeleteOps(deletes, fs.DDLTables)
	return deleteOpsToSteps(deletes, ctx)
}
