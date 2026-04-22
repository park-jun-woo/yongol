//ff:func feature=validate type=util control=sequence topic=sqlc
//ff:what toSnake — PascalCase → snake_case (에러 메시지용)

package ssac_sqlc

import "github.com/ettle/strcase"

func toSnake(s string) string {
	return strcase.ToSnake(s)
}
