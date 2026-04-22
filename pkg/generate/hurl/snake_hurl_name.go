//ff:func feature=gen-hurl type=util control=sequence
//ff:what snakeHurlName — PascalCase/kebab-case/mixed → snake_case (hurl 변수명 정규화)
package hurl

import (
	"strings"

	"github.com/ettle/strcase"
)

// snakeHurlName normalizes an identifier into snake_case suitable for a
// hurl variable name (letters / digits / underscore only).
//
// Accepts any of: PascalCase (`GigID`), camelCase (`gigId`), kebab-case
// (`audit-logs`), raw snake (`audit_log`). Hyphens are replaced with
// underscores before strcase conversion so hyphenated segments don't
// survive into hurl output.
func snakeHurlName(s string) string {
	if s == "" {
		return s
	}
	return strcase.ToSnake(strings.ReplaceAll(s, "-", "_"))
}
