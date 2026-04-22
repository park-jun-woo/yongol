//ff:func feature=validate type=util control=sequence topic=ddl-statemachine
//ff:what diagramDDLTarget — stateDiagram ID → (DDL table, column) 매핑

package ddl_statemachine

import (
	"strings"

	"github.com/jinzhu/inflection"
)

// diagramDDLTarget maps a stateDiagram ID to its corresponding DDL table and
// column. Two conventions are supported:
//   - `<PascalEntity>` → `<entity_plural>`, column = `status`
//   - `<entity>_<column>` → `<entity_plural>`, column = `<column>`
func diagramDDLTarget(id string) (string, string) {
	idx := strings.LastIndex(id, "_")
	if idx <= 0 || idx == len(id)-1 {
		return inflection.Plural(strings.ToLower(id)), "status"
	}
	entity := strings.ToLower(id[:idx])
	column := id[idx+1:]
	return inflection.Plural(entity), column
}
