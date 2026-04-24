//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what pascalToSnake — "ExecutionLog" → "execution_log"

package ssac

import (
	"strings"
)

// pascalToSnake converts "ExecutionLog" → "execution_log" for comparison
// with DDL table names. Rolls over the minimal alphabet we actually see in
// sqlc model names (ASCII, no digits leading) — for anything richer we
// defer to dedicated case converters elsewhere in the codebase.
func pascalToSnake(s string) string {
	var out strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			out.WriteByte('_')
		}
		if r >= 'A' && r <= 'Z' {
			out.WriteRune(r + ('a' - 'A'))
		} else {
			out.WriteRune(r)
		}
	}
	return out.String()
}
