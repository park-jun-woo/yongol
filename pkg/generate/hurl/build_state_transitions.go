//ff:func feature=gen-hurl type=util control=iteration dimension=2
//ff:what buildStateTransitions — stateDiagram BFS + branch-skip 순서로 전이 step 생성
package hurl

import (
	"sort"
)

// buildStateTransitions creates steps for state transitions in BFS order,
// skipping branch transitions (keeps only the first per from-state).
func buildStateTransitions(ctx *scenarioCtx) []step {
	fs := ctx.fs
	if len(fs.StateDiagrams) == 0 || fs.OpenAPIDoc == nil {
		return nil
	}
	order := buildTransitionOrder(fs.StateDiagrams)
	skip := buildBranchSkipSet(fs.StateDiagrams, order)
	opLookup := buildOpLookup(fs)

	type eventStep struct {
		event string
		ord   int
	}
	var events []eventStep
	for _, d := range fs.StateDiagrams {
		if d == nil {
			continue
		}
		for _, tr := range d.Transitions {
			if skip[tr.Event] {
				continue
			}
			if ord, ok := order[tr.Event]; ok {
				events = append(events, eventStep{event: tr.Event, ord: ord})
			}
		}
	}
	sort.SliceStable(events, func(i, j int) bool {
		if events[i].ord != events[j].ord {
			return events[i].ord < events[j].ord
		}
		return events[i].event < events[j].event
	})

	seen := map[string]bool{}
	var steps []step
	for _, es := range events {
		if seen[es.event] {
			continue
		}
		seen[es.event] = true
		s, ok := buildTransitionStep(ctx, opLookup, es.event)
		if !ok {
			continue
		}
		steps = append(steps, s)
	}
	return steps
}
