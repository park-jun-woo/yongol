//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what requestBodyFieldTypes — Operation의 requestBody 스키마에서 top-level property 이름 -> OpenAPI type 맵 추출

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// requestBodyFieldTypes extracts top-level property names and their OpenAPI
// type (e.g. "string", "object", "array") from the request body schema of an
// operation. allOf members are resolved one level deep. Properties without a
// declared type map to "".
func requestBodyFieldTypes(op *openapi3.Operation) map[string]string {
	out := make(map[string]string)
	if op == nil || op.RequestBody == nil || op.RequestBody.Value == nil {
		return out
	}
	for _, mt := range op.RequestBody.Value.Content {
		if mt == nil || mt.Schema == nil || mt.Schema.Value == nil {
			continue
		}
		collectPropTypes(out, mt.Schema.Value)
		return out
	}
	return out
}
