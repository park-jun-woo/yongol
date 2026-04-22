//ff:func feature=gen-hurl type=util control=sequence
//ff:what buildReadStep — 단일 GET path 에서 hurl step 생성 (path param 해소 실패 시 ok=false)
package hurl

// buildReadStep constructs one hurl step for a single GET path. Returns
// ok=false when either the path has no GET operation or captures cannot
// resolve the path parameters — callers skip such entries. The isFirst flag
// carries the "# ===== Reads =====" section comment onto the first step only.
func buildReadStep(ctx *scenarioCtx, path string, isFirst bool) (step, bool) {
	pathItem := ctx.fs.OpenAPIDoc.Paths.Find(path)
	if pathItem == nil || pathItem.Get == nil {
		return step{}, false
	}
	if !canResolvePathParams(path, ctx.captures) {
		return step{}, false
	}
	op := pathItem.Get
	s := step{
		Method:      "GET",
		Path:        substitutePathParams(path, ctx.captures),
		OperationID: op.OperationID,
		NeedsAuth:   needsAuth(op, ctx.fs.OpenAPIDoc),
		TokenVar:    resolveTokenVar(op.OperationID, ctx.roleMap, ctx.captures),
		StatusCode:  inferSuccessStatus(op),
		Assertions:  generateResponseAssertions(op),
	}
	if isFirst {
		s.Comment = "# ===== Reads ====="
	}
	return s, true
}
