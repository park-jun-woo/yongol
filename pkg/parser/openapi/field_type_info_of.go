//ff:func feature=openapi-parse type=parser control=sequence
//ff:what fieldTypeInfoOf — 스키마에서 primary 타입과 format을 FieldTypeInfo로 추출한다

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// fieldTypeInfoOf extracts the primary type and format from a schema.
func fieldTypeInfoOf(s *openapi3.Schema) FieldTypeInfo {
	info := FieldTypeInfo{Format: s.Format}
	if s.Type != nil {
		if types := s.Type.Slice(); len(types) > 0 {
			info.Type = types[0]
		}
	}
	return info
}
