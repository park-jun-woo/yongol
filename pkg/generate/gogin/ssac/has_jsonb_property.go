//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what hasJSONBProperty — schema 의 properties 중 JSONB shape 이 있는지 여부

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// hasJSONBProperty returns true when any property of schema is a JSONB
// shape (see isJSONBProperty). Used by the convert-file emitter to
// decide whether to import encoding/json; schemas without JSONB fields
// skip the import to avoid the Go "imported and not used" error.
func hasJSONBProperty(schema *openapi3.Schema) bool {
	if schema == nil {
		return false
	}
	for _, p := range schema.Properties {
		if isJSONBProperty(p) {
			return true
		}
	}
	return false
}
