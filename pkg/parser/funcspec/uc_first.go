//ff:func feature=funcspec type=util control=sequence
//ff:what 문자열을 Go PascalCase로 변환한다
package funcspec

import "github.com/ettle/strcase"

// ucFirst converts to Go PascalCase (uppercases the first character with Go initialism handling).
func ucFirst(s string) string {
	return strcase.ToGoPascal(s)
}
