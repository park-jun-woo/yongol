//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what snakeCase — PascalCase/camelCase → snake_case 변환

package ssac

import (
	"strings"
)

// snakeCase converts a PascalCase or camelCase identifier to snake_case.
func snakeCase(s string) string {
	if s == "" {
		return s
	}
	var b strings.Builder
	for i, r := range s {
		appendSnakeRune(&b, i, r)
	}
	return b.String()
}
