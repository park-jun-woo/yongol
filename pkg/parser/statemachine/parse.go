//ff:func feature=statemachine type=parser control=iteration dimension=1 topic=states
//ff:what Mermaid stateDiagram 텍스트를 파싱하여 StateDiagram 구조체를 반환한다
package statemachine

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"github.com/ettle/strcase"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

var (
	// [*] --> stateName
	reInitial = regexp.MustCompile(`\[\*\]\s*-->\s*(\w+)`)
	// stateA --> stateB: EventName
	reTransition = regexp.MustCompile(`(\w+)\s*-->\s*(\w+)\s*:\s*(\w+)`)
)

// Parse parses Mermaid stateDiagram content with a given ID.
// file is the source file path used in diagnostics.
func Parse(id, content, file string) (*StateDiagram, []diagnostic.Diagnostic) {
	// Extract content inside ```mermaid ... ``` block.
	mermaidContent, mermaidStartLine := extractMermaidBlockWithLine(content)
	if mermaidContent == "" {
		return nil, []diagnostic.Diagnostic{{
			File:    file,
			Line:    1,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("no mermaid stateDiagram block found in %s", id),
		}}
	}

	// Symbol is PascalCase of ID. Computed once at parse time so all
	// downstream emitters (render_state_file, render_can_transition_file,
	// render_next_state_file, ssac/build_state) use a single source of
	// truth for the exported Go identifier. Writing Symbol here instead
	// of at each emit site removes the foot-gun where a new emitter
	// forgets the case conversion and silently produces an unexported
	// identifier that fails at go build.
	d := &StateDiagram{ID: id, Symbol: strcase.ToPascal(id), File: file}
	var diags []diagnostic.Diagnostic

	// Parse line by line to track line numbers.
	lines := strings.Split(mermaidContent, "\n")
	stateSet := make(map[string]bool)

	for i, line := range lines {
		// mermaidStartLine is the 1-based line of ```mermaid.
		// mermaidContent starts with the newline after ```mermaid, so lines[0]
		// corresponds to that ```mermaid line itself; lines[i] is at line
		// mermaidStartLine + i.
		lineNum := mermaidStartLine + i

		// Parse initial state.
		if m := reInitial.FindStringSubmatch(line); len(m) > 1 && d.InitialState == "" {
			d.InitialState = m[1]
			stateSet[m[1]] = true
		}

		// Parse transition.
		if m := reTransition.FindStringSubmatch(line); len(m) > 3 {
			from, to, event := m[1], m[2], m[3]
			d.Transitions = append(d.Transitions, Transition{
				From:  from,
				To:    to,
				Event: event,
				Line:  lineNum,
			})
			stateSet[from] = true
			stateSet[to] = true
		}
	}

	diags = append(diags, checkCaseConflicts(id, file, stateSet, lines, mermaidStartLine)...)

	for s := range stateSet {
		d.States = append(d.States, s)
	}
	sort.Strings(d.States)

	if len(d.Transitions) == 0 {
		diags = append(diags, diagnostic.Diagnostic{
			File:    file,
			Line:    mermaidStartLine,
			Phase:   diagnostic.PhaseParse,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("no transitions found in stateDiagram %s", id),
		})
	}

	if len(diags) > 0 {
		return nil, diags
	}
	return d, nil
}
