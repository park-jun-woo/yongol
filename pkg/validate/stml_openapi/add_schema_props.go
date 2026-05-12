//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what addSchemaProps — 스키마 top-level + allOf 프로퍼티를 responseFieldInfo 맵에 추가

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// addSchemaProps adds top-level properties and resolves allOf one level.
func addSchemaProps(out map[string]responseFieldInfo, s *openapi3.Schema) {
	if s == nil {
		return
	}
	for name, ref := range s.Properties {
		out[name] = responseFieldInfo{typ: schemaType(ref)}
	}
	for _, allOfRef := range s.AllOf {
		if allOfRef == nil || allOfRef.Value == nil {
			continue
		}
		for name, ref := range allOfRef.Value.Properties {
			out[name] = responseFieldInfo{typ: schemaType(ref)}
		}
	}
}
