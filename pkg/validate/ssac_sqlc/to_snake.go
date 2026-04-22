//ff:func feature=validate type=util control=sequence topic=sqlc
//ff:what toSnake — convert PascalCase to snake_case (used in error messages)

package ssac_sqlc

import "github.com/ettle/strcase"

func toSnake(s string) string {
	return strcase.ToSnake(s)
}
