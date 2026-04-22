//ff:func feature=gen-hurl type=util control=sequence
//ff:what buildTransitionStep — 단일 state transition event 에서 hurl step 조립
package hurl

// buildTransitionStep constructs one hurl step for a single stateDiagram
// event. Returns ok=false when the event has no mapped OpenAPI operation,
// the operation cannot be located at the resolved path, or captures cannot
// resolve its path parameters — callers skip such events.
func buildTransitionStep(ctx *scenarioCtx, opLookup map[string]opInfo, event string) (step, bool) {
	info, ok := opLookup[event]
	if !ok {
		return step{}, false
	}
	// info.Method comes from buildOpLookup in upper case (e.g. "POST"); PathItem.Operations()
	// keys are http.MethodX constants which are already upper case. No conversion needed.
	op := findOperation(ctx.fs.OpenAPIDoc, info.Path, info.Method)
	if op == nil {
		return step{}, false
	}
	if !canResolvePathParams(info.Path, ctx.captures) {
		return step{}, false
	}
	body := generateRequestBody(op, ctx.fs, "")
	s := step{
		Method:      info.Method,
		Path:        substitutePathParams(info.Path, ctx.captures),
		OperationID: event,
		NeedsAuth:   needsAuth(op, ctx.fs.OpenAPIDoc),
		TokenVar:    resolveTokenVar(event, ctx.roleMap, ctx.captures),
		HasBody:     body != "",
		BodyJSON:    body,
		StatusCode:  inferSuccessStatus(op),
	}
	return s, true
}
