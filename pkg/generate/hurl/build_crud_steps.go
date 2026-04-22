//ff:func feature=gen-hurl type=util control=sequence
//ff:what buildCRUDSteps — Phase3 CRUD+StateTransitions: Create/Transitions/Update 순서 (Read는 Phase4)
package hurl

// buildCRUDSteps produces Phase 3 steps: Create, State Transitions, Update.
// GET endpoints are emitted separately by buildReads in Phase 4.
func buildCRUDSteps(ctx *scenarioCtx) []step {
	if ctx.fs.OpenAPIDoc == nil {
		return nil
	}
	var all []step
	all = append(all, buildCreateSteps(ctx)...)
	all = append(all, buildStateTransitions(ctx)...)
	all = append(all, buildUpdateSteps(ctx)...)
	return all
}
