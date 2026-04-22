//ff:func feature=validate type=util control=iteration dimension=1 topic=scenario-check
//ff:what collectResponseCodes — OpenAPI Operation에서 응답 상태코드 수집

package openapi_hurl

import "github.com/getkin/kin-openapi/openapi3"

// collectResponseCodes collects response codes from an operation.
func collectResponseCodes(op *openapi3.Operation) map[string]bool {
	codes := make(map[string]bool)
	if op == nil || op.Responses == nil {
		return codes
	}
	for code := range op.Responses.Map() {
		codes[code] = true
	}
	return codes
}
