//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what requestBodyFields — Operation의 requestBody 스키마에서 top-level property 이름 집합 추출

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// requestBodyFields extracts top-level property names from the request body
// schema of an operation. allOf members are resolved one level.
func requestBodyFields(op *openapi3.Operation) map[string]struct{} {
	out := make(map[string]struct{})
	if op == nil || op.RequestBody == nil || op.RequestBody.Value == nil {
		return out
	}
	for _, mt := range op.RequestBody.Value.Content {
		if mt == nil || mt.Schema == nil || mt.Schema.Value == nil {
			continue
		}
		collectPropNames(out, mt.Schema.Value)
		return out
	}
	return out
}

// collectPropNames collects property names from schema + allOf.
func collectPropNames(out map[string]struct{}, s *openapi3.Schema) {
	if s == nil {
		return
	}
	for name := range s.Properties {
		out[name] = struct{}{}
	}
	for _, allOfRef := range s.AllOf {
		if allOfRef == nil || allOfRef.Value == nil {
			continue
		}
		for name := range allOfRef.Value.Properties {
			out[name] = struct{}{}
		}
	}
}
