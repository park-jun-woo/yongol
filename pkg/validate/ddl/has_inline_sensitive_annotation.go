//ff:func feature=validate type=util control=sequence topic=ddl-structural
//ff:what hasInlineSensitiveAnnotation — report whether a DDL line carries a -- @sensitive / -- @nosensitive annotation
package ddl

import "strings"

// hasInlineSensitiveAnnotation reports whether the given DDL line carries
// `-- @sensitive` (case-insensitive). `-- @nosensitive` suppresses the
// warning and is treated as annotated.
func hasInlineSensitiveAnnotation(line string) bool {
	lower := strings.ToLower(line)
	return strings.Contains(lower, "-- @sensitive") ||
		strings.Contains(lower, "--@sensitive") ||
		strings.Contains(lower, "-- @nosensitive") ||
		strings.Contains(lower, "--@nosensitive")
}
