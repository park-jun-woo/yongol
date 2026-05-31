//ff:func feature=gen-nestjs type=util control=sequence
//ff:what tsSourceExpr — Go 스타일 dotted access 를 camelCase TypeScript 표현식으로 변환

package ssac

import "strings"

// tsSourceExpr converts a Go-style dotted access to camelCase TypeScript.
// "token.AccessToken" → "token.accessToken", "user" → "user".
func tsSourceExpr(source string) string {
	dotIdx := strings.Index(source, ".")
	if dotIdx < 0 {
		return source
	}
	obj := source[:dotIdx]
	field := source[dotIdx+1:]
	return obj + "." + lcFirst(field)
}
