//ff:func feature=validate type=util control=sequence topic=ddl-statemachine
//ff:what checkStateFieldColumn — verifies that the DDL column referenced by a single stateDiagram exists

package ddl_statemachine

import (
	"strings"

	"github.com/jinzhu/inflection"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// checkStateFieldColumn returns a diagnostic when a stateDiagram shaped
// `<entity>_<field>` does not correspond to an existing DDL column. Returns
// nil when the diagram is skipped (no underscore) or the column is present.
func checkStateFieldColumn(diagramID string, g *rule.Ground) *diagnostic.Diagnostic {
	idx := strings.LastIndex(diagramID, "_")
	if idx <= 0 || idx == len(diagramID)-1 {
		return nil
	}
	entity := diagramID[:idx]
	column := diagramID[idx+1:]
	table := inflection.Plural(entity)

	cols := g.Lookup["DDL.column."+table]
	if cols != nil && cols[column] {
		return nil
	}
	return &diagnostic.Diagnostic{
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[XDM-27] state diagram " + diagramID + " expects DDL column " + table + "." + column + " which does not exist",
		Advice:  "Add column " + column + " to DDL table " + table + ", or rename the stateDiagram",
	}
}
