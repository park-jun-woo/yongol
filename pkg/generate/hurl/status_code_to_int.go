//ff:func feature=gen-hurl type=util control=selection
//ff:what statusCodeToInt — HTTP 상태 코드 문자열을 정수로 변환
package hurl

// statusCodeToInt converts a status code string to an integer.
func statusCodeToInt(code string) int {
	switch code {
	case "200":
		return 200
	case "201":
		return 201
	case "204":
		return 204
	default:
		return 200
	}
}
