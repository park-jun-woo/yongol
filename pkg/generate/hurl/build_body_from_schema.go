//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what buildBodyFromSchema — OpenAPI schema properties에서 dummy value map 생성
package hurl

import (
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildBodyFromSchema generates a map from an OpenAPI schema properties.
func buildBodyFromSchema(schema *openapi3.Schema, fs *yongol.Fullstack, role string) map[string]any {
	if schema == nil {
		return nil
	}
	result := map[string]any{}
	for name, propRef := range schema.Properties {
		if propRef == nil || propRef.Value == nil {
			continue
		}
		result[name] = resolveDummyValue(name, propRef.Value, fs, role)
	}
	if len(result) == 0 {
		return nil
	}
	return result
}
