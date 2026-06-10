//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what opDeclaresErrorResponse — operation 응답 맵에 4xx/5xx 상태 코드 키가 1개 이상 선언됐는지 판정

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// opDeclaresErrorResponse reports whether the operation declares at least
// one 4xx/5xx response. Status keys are matched by their first character,
// which covers both concrete codes ("400", "503") and range wildcards
// ("4XX", "5XX"); the "default" key is excluded — TM-29 only cares about
// explicitly declared error responses, regardless of body schema shape.
func opDeclaresErrorResponse(op *openapi3.Operation) bool {
	if op == nil || op.Responses == nil {
		return false
	}
	for code := range op.Responses.Map() {
		if len(code) > 0 && (code[0] == '4' || code[0] == '5') {
			return true
		}
	}
	return false
}
