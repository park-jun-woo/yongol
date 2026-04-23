//ff:func feature=gen-hurl type=util control=iteration dimension=2
//ff:what buildStateTransitions — Phase004: 모든 state transition emit + bridge 재활성 삽입

package hurl

// buildStateTransitions creates steps for every state transition in
// order, inserting bridging events when a terminal transition requires
// a prerequisite state the linear walk has already left.
//
// BUG-016 / Phase004 — the previous implementation kept only the first
// transition per from-state via buildBranchSkipSet, dropping events
// like ExecuteWorkflow and ArchiveWorkflow from smoke. The new walker
// emits all transitions. When the simulated current state does not
// match the event's from-state (e.g. after PauseWorkflow the state is
// `paused`, but ArchiveWorkflow needs `active`), a bridging transition
// is inserted first (paused → active via ActivateWorkflow) so the
// smoke sequence stays executable end-to-end.
func buildStateTransitions(ctx *scenarioCtx) []step {
	fs := ctx.fs
	if len(fs.StateDiagrams) == 0 || fs.OpenAPIDoc == nil {
		return nil
	}
	order := buildTransitionOrder(fs.StateDiagrams)
	opLookup := buildOpLookup(fs)
	events := orderedStateEvents(fs.StateDiagrams, order)

	var steps []step
	state := initialState(fs.StateDiagrams)
	for _, ev := range events {
		bridges, nextState := bridgeSteps(ctx, fs.StateDiagrams, opLookup, state, ev.event)
		state = nextState
		steps = append(steps, bridges...)
		s, ok := buildTransitionStep(ctx, opLookup, ev.event)
		if !ok {
			continue
		}
		steps = append(steps, s)
		state = eventTargetState(fs.StateDiagrams, state, ev.event)
	}
	return steps
}
