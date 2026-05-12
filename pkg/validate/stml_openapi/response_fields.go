//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what responseFields — Operation의 200 응답 스키마에서 top-level property -> type 맵 추출

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// responseFields extracts top-level property names and their types from
// the first 2xx response schema of the given operation. allOf members
// are resolved one level deep.
func responseFields(op *openapi3.Operation) map[string]responseFieldInfo {
	out := make(map[string]responseFieldInfo)
	if op == nil || op.Responses == nil {
		return out
	}
	for _, code := range []string{"200", "201"} {
		resp := op.Responses.Status(statusInt(code))
		if resp == nil || resp.Value == nil {
			continue
		}
		collectResponseProps(out, resp.Value)
		if len(out) > 0 {
			return out
		}
	}
	return out
}
