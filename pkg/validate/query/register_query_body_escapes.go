//ff:func feature=validate type=util control=iteration dimension=1 topic=query-structural
//ff:what registerQueryBodyEscapes — `-- +<name>` / `-- @<name>` escape 마커를 body.Escapes 에 기록

package query

import (
	"strings"
)

// escapeMarkers lists all supported query escape-hatch names.
var escapeMarkers = []string{"no-pagination", "allow-truncate", "allow-sensitive", "skip-column-check"}

// registerQueryBodyEscapes scans `trimmed` (a SQL comment line) for every
// supported escape marker (`+<name>` or `@<name>` form) and records them
// into body.Escapes. Sets body.HasStop when any marker matches.
func registerQueryBodyEscapes(body *queryBody, trimmed string) {
	for _, esc := range escapeMarkers {
		if strings.Contains(trimmed, "+"+esc) || strings.Contains(trimmed, "@"+esc) {
			body.Escapes["@"+esc] = true
			body.HasStop = true
		}
	}
}
