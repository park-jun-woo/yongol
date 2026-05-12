//ff:func feature=validate type=util control=selection topic=stml-openapi
//ff:what statusInt — 2xx 상태 코드 문자열을 정수로 변환

package stml_openapi

// statusInt converts a status code string to int. Only used for 2xx.
func statusInt(code string) int {
	switch code {
	case "200":
		return 200
	case "201":
		return 201
	default:
		return 0
	}
}
