//ff:func feature=gen-nestjs type=util control=selection
//ff:what httpStatusConst — HTTP 상태코드 → NestJS HttpStatus enum 이름 변환

package ssac

import "fmt"

// httpStatusConst maps an HTTP status code to the NestJS HttpStatus enum name.
func httpStatusConst(code int) string {
	switch code {
	case 400:
		return "BAD_REQUEST"
	case 401:
		return "UNAUTHORIZED"
	case 403:
		return "FORBIDDEN"
	case 404:
		return "NOT_FOUND"
	case 409:
		return "CONFLICT"
	case 422:
		return "UNPROCESSABLE_ENTITY"
	case 500:
		return "INTERNAL_SERVER_ERROR"
	default:
		return fmt.Sprintf("/* %d */BAD_REQUEST", code)
	}
}
