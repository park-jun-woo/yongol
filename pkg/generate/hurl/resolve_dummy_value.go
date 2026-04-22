//ff:func feature=gen-hurl type=util control=sequence
//ff:what resolveDummyValue — 필드명+schema로 적절한 dummy 값 결정 (role/enum/type 순)
package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// resolveDummyValue decides the dummy value for a field.
// Priority: role override → DDL CHECK enum → OpenAPI enum → field-name hint → type default.
func resolveDummyValue(name string, prop *openapi3.Schema, fs *yongol.Fullstack, role string) any {
	if name == "role" && role != "" {
		return role
	}
	if enumVal := lookupCheckEnum(fs, name); enumVal != "" {
		return enumVal
	}
	if len(prop.Enum) > 0 {
		return prop.Enum[0]
	}
	if v := dummyFieldHint(name); v != nil {
		return v
	}
	return generateDummyValue(prop.Type, prop.Format)
}
