//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-openapi
//ff:what hasExplicit2xx — Operation에 명시적 2xx 응답이 있는지 확인

package openapi_ssac

import "github.com/getkin/kin-openapi/openapi3"

// hasExplicit2xx returns true if op.Responses includes an explicit 2xx code (e.g. "200").
func hasExplicit2xx(op *openapi3.Operation) bool {
	if op == nil || op.Responses == nil {
		return false
	}
	for code := range op.Responses.Map() {
		if len(code) == 3 && code[0] == '2' {
			return true
		}
	}
	return false
}
