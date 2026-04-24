//ff:func feature=validate type=util control=sequence topic=states
//ff:what isStatefulResource — path segment → (table, state column, diagram) stateful 리소스 판정

package ssac_statemachine

import (
	"strings"

	"github.com/jinzhu/inflection"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// statefulTarget collects identifiers XSM-27 needs when composing its advice.
type statefulTarget struct {
	Resource    string                   // singular lowercase resource name (e.g. "workflow")
	Table       string                   // plural lowercase DDL table name (e.g. "workflows")
	Diagram     *statemachine.StateDiagram
	StateColumn string                   // DDL column that holds the state (default "status")
	Model       string                   // PascalCase model name (e.g. "Workflow")
}

// isStatefulResource determines whether the resource named by the first path
// segment of an OpenAPI path is "stateful" — i.e. there exists a Mermaid
// stateDiagram for it AND the corresponding DDL column's DEFAULT matches the
// diagram's `[*] --> X` initial state (XDM-28 linkage).
//
// Two diagram-ID conventions are supported (mirrors diagramDDLTarget in
// ddl_statemachine): `<resource>` → table=plural(resource), column=status;
// and `<entity>_<column>` → table=plural(entity), column=<column>.
//
// Returns the matched target when stateful, or nil when any of the preconditions fail.
func isStatefulResource(path string, diagrams []*statemachine.StateDiagram, g *rule.Ground) *statefulTarget {
	resource := firstPathSegment(path)
	if resource == "" {
		return nil
	}
	// Path segments are typically plural ("/workflows/{id}"); the diagram ID
	// is typically the singular ("workflow"). Use inflection.Singular to
	// normalise, while still accepting paths that are already singular.
	singular := strings.ToLower(inflection.Singular(resource))
	plural := strings.ToLower(inflection.Plural(singular))

	for _, d := range diagrams {
		if d == nil {
			continue
		}
		diagramID := d.ID
		table, column := diagramIDToTable(diagramID)
		if table != plural {
			continue
		}
		if d.InitialState == "" {
			continue
		}
		if g == nil {
			continue
		}
		got := g.Types["DDL.default.value."+table+"."+column]
		if got == "" || got != d.InitialState {
			continue
		}
		return &statefulTarget{
			Resource:    singular,
			Table:       table,
			Diagram:     d,
			StateColumn: column,
			Model:       pascalCaseFromLower(singular),
		}
	}
	return nil
}

// firstPathSegment returns the first non-empty segment of a URL path, skipping
// leading slashes and curly-brace parameters.
func firstPathSegment(path string) string {
	for _, part := range strings.Split(path, "/") {
		if part == "" {
			continue
		}
		if strings.HasPrefix(part, "{") {
			continue
		}
		return part
	}
	return ""
}

// diagramIDToTable mirrors ddl_statemachine.diagramDDLTarget so XSM-27 can
// identify the DDL table/column a diagram describes without importing the
// sibling validate package (that would create a cycle through yongol).
func diagramIDToTable(id string) (string, string) {
	idx := strings.LastIndex(id, "_")
	if idx <= 0 || idx == len(id)-1 {
		return inflection.Plural(strings.ToLower(id)), "status"
	}
	entity := strings.ToLower(id[:idx])
	column := id[idx+1:]
	return inflection.Plural(entity), column
}

// pascalCaseFromLower returns "workflow" → "Workflow" (single-word resources
// only; multi-word diagram IDs already carry their own Symbol on StateDiagram).
func pascalCaseFromLower(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
