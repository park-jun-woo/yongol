//ff:func feature=validate type=util control=iteration dimension=1 topic=states
//ff:what isStatefulResource — path segment → (table, state column, diagram) stateful 리소스 판정

package ssac_statemachine

import (
	"strings"

	"github.com/jinzhu/inflection"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

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
		if target := matchStatefulDiagram(d, plural, singular, g); target != nil {
			return target
		}
	}
	return nil
}
