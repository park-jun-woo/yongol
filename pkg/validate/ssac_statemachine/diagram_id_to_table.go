//ff:func feature=validate type=util control=sequence topic=states
//ff:what diagramIDToTable — state diagram ID → (table, state column) 쌍 도출

package ssac_statemachine

import (
	"strings"

	"github.com/jinzhu/inflection"
)

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
