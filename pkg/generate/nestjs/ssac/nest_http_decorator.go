//ff:func feature=gen-nestjs type=util control=selection
//ff:what nestHTTPDecorator — HTTP 메서드 문자열 → NestJS 데코레이터 이름 변환

package ssac

import "strings"

// nestHTTPDecorator maps an HTTP method string to a NestJS decorator name.
func nestHTTPDecorator(method string) string {
	switch strings.ToUpper(method) {
	case "GET":
		return "Get"
	case "POST":
		return "Post"
	case "PUT":
		return "Put"
	case "DELETE":
		return "Delete"
	case "PATCH":
		return "Patch"
	default:
		return "Get"
	}
}
