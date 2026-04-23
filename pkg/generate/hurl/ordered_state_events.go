//ff:func feature=gen-hurl type=util control=iteration dimension=2
//ff:what orderedStateEvents — diagram transitions 을 order 맵 기준으로 정렬·중복 제거

package hurl

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// orderedStateEvents returns the dedup'd ordered list of state events
// the smoke walker consumes. The returned slice contains each event
// exactly once, sorted by its index in `order` (ties broken by event
// name for stability).
func orderedStateEvents(diagrams []*statemachine.StateDiagram, order map[string]int) []orderedStateEvent {
	var events []orderedStateEvent
	seen := map[string]bool{}
	for _, d := range diagrams {
		if d == nil {
			continue
		}
		for _, tr := range d.Transitions {
			if seen[tr.Event] {
				continue
			}
			ord, ok := order[tr.Event]
			if !ok {
				continue
			}
			seen[tr.Event] = true
			events = append(events, orderedStateEvent{event: tr.Event, ord: ord})
		}
	}
	sort.SliceStable(events, func(i, j int) bool {
		if events[i].ord != events[j].ord {
			return events[i].ord < events[j].ord
		}
		return events[i].event < events[j].event
	})
	return events
}
