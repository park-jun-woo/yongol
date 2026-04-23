//ff:func feature=util type=util control=sequence topic=string-convert
//ff:what PascalToSnake — PascalCase / camelCase → snake_case (ettle/strcase 경유)

package caseconv

import "github.com/ettle/strcase"

// PascalToSnake converts PascalCase / camelCase to snake_case via ettle/strcase.
func PascalToSnake(s string) string {
	return strcase.ToSnake(s)
}
