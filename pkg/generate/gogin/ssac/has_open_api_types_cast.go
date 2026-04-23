//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what hasOpenAPITypesCast — schema properties 중 openapi_types.Email/UUID 캐스트가 필요한 필드가 있는지

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// hasOpenAPITypesCast returns true when at least one of schema's string
// properties has a format that apiCastFor maps to openapi_types.*.
// Drives the conditional import of github.com/oapi-codegen/runtime/types
// in the generated convert file.
func hasOpenAPITypesCast(schema *openapi3.Schema) bool {
	if schema == nil {
		return false
	}
	for jsonName, ref := range schema.Properties {
		cast := apiCastFor("", jsonName, ref)
		if cast == "openapi_types.Email" || cast == "openapi_types.UUID" {
			return true
		}
	}
	return false
}
