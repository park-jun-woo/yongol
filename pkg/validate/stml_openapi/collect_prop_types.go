//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what collectPropTypes — 스키마 + allOf 에서 property 이름 -> 첫 번째 선언 type 을 수집

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// collectPropTypes records each property's first declared type from schema +
// allOf. Properties without a declared type map to "".
func collectPropTypes(out map[string]string, s *openapi3.Schema) {
	if s == nil {
		return
	}
	recordPropTypes(out, s.Properties)
	for _, allOfRef := range s.AllOf {
		if allOfRef == nil || allOfRef.Value == nil {
			continue
		}
		recordPropTypes(out, allOfRef.Value.Properties)
	}
}
