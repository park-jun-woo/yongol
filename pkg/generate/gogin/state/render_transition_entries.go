//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what renderTransitionEntries — transitionMap의 각 상태 엔트리를 Go map 리터럴로 렌더링

package state

import (
	"fmt"
	"strings"
)

// renderTransitionEntries writes sorted state→event→nextState entries into the
// builder as Go map literal lines.
func renderTransitionEntries(b *strings.Builder, transMap map[string]map[string]string) {
	states := sortedKeys(transMap)
	for _, st := range states {
		events := transMap[st]
		evNames := sortedKeys(events)
		if len(evNames) == 1 {
			fmt.Fprintf(b, "\t%q: {%q: %q},\n", st, evNames[0], events[evNames[0]])
			continue
		}
		fmt.Fprintf(b, "\t%q: {\n", st)
		for _, ev := range evNames {
			fmt.Fprintf(b, "\t\t%q: %q,\n", ev, events[ev])
		}
		b.WriteString("\t},\n")
	}
}
