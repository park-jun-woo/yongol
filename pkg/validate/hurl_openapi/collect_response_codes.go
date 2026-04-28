//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what collectResponseCodes — operation 의 응답 코드 키 집합 생성

package hurl_openapi

import "github.com/getkin/kin-openapi/openapi3"

// collectResponseCodes collects the response-code keys declared on an
// operation. Empty op / absent responses yield an empty map so callers
// can use `codes[key]` guards without nil checks.
func collectResponseCodes(op *openapi3.Operation) map[string]bool {
	codes := map[string]bool{}
	if op == nil || op.Responses == nil {
		return codes
	}
	for code := range op.Responses.Map() {
		codes[code] = true
	}
	return codes
}
