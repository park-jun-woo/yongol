//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what collectDeleteOps — OpenAPI에서 DELETE operation 목록 수집 (shape-detected auth 제외)
package hurl

// collectDeleteOps collects all DELETE operations from OpenAPI,
// excluding ops classified as signup/login by ctx.authOpIDs.
func collectDeleteOps(ctx *scenarioCtx) []deleteOp {
	var ops []deleteOp
	for path, pathItem := range ctx.fs.OpenAPIDoc.Paths.Map() {
		if pathItem.Delete == nil || pathItem.Delete.OperationID == "" {
			continue
		}
		if isAuthOpID(ctx, pathItem.Delete.OperationID) {
			continue
		}
		ops = append(ops, deleteOp{path: path, opID: pathItem.Delete.OperationID, resource: resourceFromPath(path)})
	}
	return ops
}
