//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what collectRequiredNames — 스키마 + allOf 의 required 필드 이름 집합 수집

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// collectRequiredNames collects the required property names from schema + allOf
// (one level), the same traversal collectPropNames uses for property names.
func collectRequiredNames(s *openapi3.Schema) map[string]struct{} {
	out := make(map[string]struct{})
	if s == nil {
		return out
	}
	for _, name := range s.Required {
		out[name] = struct{}{}
	}
	for _, allOfRef := range s.AllOf {
		if allOfRef == nil || allOfRef.Value == nil {
			continue
		}
		for _, name := range allOfRef.Value.Required {
			out[name] = struct{}{}
		}
	}
	return out
}
