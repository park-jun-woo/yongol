//ff:func feature=validate type=util control=sequence topic=ddl-structural
//ff:what hasInlineSensitiveAnnotation — DDL 라인에 -- @sensitive / -- @nosensitive 어노테이션 여부
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
