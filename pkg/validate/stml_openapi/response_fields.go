//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what responseFields — Operation의 성공 2xx(200→201 중 props 있는 첫 코드) 응답 스키마에서 top-level property -> type 맵 추출

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// responseFields extracts top-level property names and their types from the
// first 2xx response schema with props, scanned in 200 → 201 priority. allOf
// members are resolved one level deep.
//
// This priority is the single source of truth for "what the frontend can read
// from a success response" and must stay in lockstep with the codegen Res<K>
// helper (pkg/generate/react/write_req_res_types.go, same 200 → 201 → void
// order). A 204/no-body op yields the empty map here and `void` there — fixing
// only one side reopens BUG-128, so change both together.
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
