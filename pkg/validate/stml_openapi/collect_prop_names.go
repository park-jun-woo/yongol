//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what collectPropNames — 스키마 + allOf 에서 property 이름을 수집

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

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
