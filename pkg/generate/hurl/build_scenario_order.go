//ff:func feature=gen-hurl type=util control=sequence
//ff:what buildScenarioOrder — 5-Phase 시나리오 순서 결정 (Auth/Prereq/CRUD/Reads/Deletes)
package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildScenarioOrder assembles all steps in 5-Phase order, threading shared
// captures + roleMap through a scenarioCtx.
func buildScenarioOrder(fs *yongol.Fullstack) []step {
	ctx := newScenarioCtx(fs)
	var steps []step
	steps = append(steps, buildAuthSteps(ctx)...)
	steps = append(steps, buildPrereqSteps(ctx)...)
	steps = append(steps, buildCRUDSteps(ctx)...)
	steps = append(steps, buildReads(ctx)...)
	steps = append(steps, buildDeleteSteps(ctx)...)
	return steps
}
