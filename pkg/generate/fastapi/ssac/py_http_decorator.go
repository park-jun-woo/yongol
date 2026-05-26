//ff:func feature=gen-fastapi type=util control=selection
//ff:what pyHTTPDecorator — HTTP 메서드 문자열 → FastAPI 라우터 데코레이터 이름 변환

package ssac

import "strings"

// pyHTTPDecorator maps an HTTP method string to a FastAPI router decorator name.
func pyHTTPDecorator(method string) string {
	switch strings.ToUpper(method) {
	case "GET":
		return "get"
	case "POST":
		return "post"
	case "PUT":
		return "put"
	case "DELETE":
		return "delete"
	case "PATCH":
		return "patch"
	default:
		return "get"
	}
}
