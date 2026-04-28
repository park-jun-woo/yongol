//ff:func feature=manifest type=util control=sequence
//ff:what GoTypeOf — Column 의 Go 타입 투영 (RawType → Go 타입)
package ddl

import "strings"

// GoTypeOf returns the Go-type projection of a parsed Column. This preserves
// backward compatibility for consumers that previously read a Go-type string
// from Table.Columns map. It strips parameterised types like VARCHAR(255)
// down to their head token before looking up the projection.
func GoTypeOf(col Column) string {
	t := col.RawType
	if idx := strings.Index(t, "("); idx > 0 {
		t = t[:idx]
	}
	t = strings.TrimSpace(t)
	return pgTypeToGo(strings.ToUpper(t))
}
