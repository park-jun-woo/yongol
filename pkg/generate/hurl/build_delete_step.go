//ff:func feature=gen-hurl type=util control=sequence
//ff:what buildDeleteStep — 단일 deleteOp 에서 hurl step 조립 (해소 실패 시 ok=false)
package hurl

// buildDeleteStep constructs one hurl step for a single deleteOp. Returns
// ok=false when either the OpenAPI path has no DELETE operation or captures
// cannot resolve the path parameters — callers skip such entries. The
// isFirst flag attaches the "# ===== Cleanup (Deletes) =====" section
// comment onto the first emitted step only.
func buildDeleteStep(ctx *scenarioCtx, d deleteOp, isFirst bool) (step, bool) {
	pathItem := ctx.fs.OpenAPIDoc.Paths.Find(d.path)
	if pathItem == nil || pathItem.Delete == nil {
		return step{}, false
	}
	if !canResolvePathParams(d.path, ctx.captures) {
		return step{}, false
	}
	op := pathItem.Delete
	s := step{
		Method:      "DELETE",
		Path:        substitutePathParams(d.path, ctx.captures),
		OperationID: d.opID,
		NeedsAuth:   needsAuth(op, ctx.fs.OpenAPIDoc),
		TokenVar:    resolveTokenVar(d.opID, ctx.roleMap, ctx.captures),
		StatusCode:  inferSuccessStatus(op),
	}
	if isFirst {
		s.Comment = "# ===== Cleanup (Deletes) ====="
	}
	return s, true
}
