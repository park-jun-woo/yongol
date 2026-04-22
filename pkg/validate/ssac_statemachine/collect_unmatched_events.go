//ff:func feature=validate type=util control=iteration dimension=1 topic=states
//ff:what diagram 이벤트 중 SSaC 함수가 없는 것 수집

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// collectUnmatchedEvents returns diagnostics for transition events that have
// no matching SSaC function.
func collectUnmatchedEvents(d *statemachine.StateDiagram, funcNames map[string]bool) []diagnostic.Diagnostic {
	if d == nil {
		return nil
	}
	file := d.File
	if file == "" {
		file = "states/" + d.ID + ".md"
	}
	// First-occurrence line per event.
	eventLine := make(map[string]int, len(d.Transitions))
	var events []string
	for _, t := range d.Transitions {
		if _, ok := eventLine[t.Event]; ok {
			continue
		}
		eventLine[t.Event] = t.Line
		events = append(events, t.Event)
	}
	var diags []diagnostic.Diagnostic
	for _, event := range events {
		if funcNames[event] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    file,
			Line:    eventLine[event],
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[XSM-23] transition event \"" + event + "\" in diagram \"" + d.ID + "\" has no matching SSaC function",
			Advice:  "stateDiagram event '" + event + "' 를 사용하는 SSaC 함수를 추가하세요",
		})
	}
	return diags
}
