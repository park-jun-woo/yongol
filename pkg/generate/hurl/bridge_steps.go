//ff:func feature=gen-hurl type=util control=sequence
//ff:what bridgeSteps — 현재 상태에서 event 실행이 불가능하면 선행 전이 step 삽입

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// bridgeSteps returns any transition steps that must precede `event` so
// that the simulated walker is in a valid from-state when `event` runs.
// Returns the steps plus the resulting state after the bridges (caller
// treats that as the new current state before running `event` itself).
//
// Used when a terminal transition (e.g. ArchiveWorkflow: active →
// archived) follows a non-terminal transition that leaves the prereq
// state (e.g. PauseWorkflow put the workflow in `paused`). A one-hop
// bridge via an event that transitions paused → active is emitted so
// Archive resolves cleanly. If no single-hop bridge exists the current
// state is returned unchanged and the caller proceeds — the attempt
// will likely 4xx at runtime, surfacing the missing transition rather
// than silently hiding it.
func bridgeSteps(ctx *scenarioCtx, diagrams []*statemachine.StateDiagram, opLookup map[string]opInfo, currentState, event string) ([]step, string) {
	fromStates := eventFromStates(diagrams, event)
	if len(fromStates) == 0 {
		return nil, currentState
	}
	if fromStates[currentState] {
		return nil, currentState
	}
	bridgeEvent, bridgeTo := findBridgeEvent(diagrams, currentState, fromStates)
	if bridgeEvent == "" {
		return nil, currentState
	}
	s, ok := buildTransitionStep(ctx, opLookup, bridgeEvent)
	if !ok {
		return nil, currentState
	}
	return []step{s}, bridgeTo
}
