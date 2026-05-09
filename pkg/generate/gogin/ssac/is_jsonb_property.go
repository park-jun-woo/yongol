//ff:func feature=gen-gogin type=util control=sequence
//ff:what isJSONBProperty — OpenAPI property schema 가 JSONB (object + additionalProperties open) 인지

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// isJSONBProperty returns true for OpenAPI property schemas that
// oapi-codegen emits as map[string]interface{} — namely `type: object`
// with no fixed properties and additionalProperties open. The sqlc side
// stores these as json.RawMessage for a JSONB column; writeConvertFunc
// bridges them via json.Unmarshal.
func isJSONBProperty(ref *openapi3.SchemaRef) bool {
	if ref == nil || ref.Value == nil {
		return false
	}
	s := ref.Value
	if s.Type == nil || !s.Type.Is("object") {
		return false
	}
	if len(s.Properties) > 0 {
		return false
	}
	// Accept either additionalProperties: true (Has) or a permissive
	// {} schema (Schema present but no constraints). oapi-codegen maps
	// both to map[string]interface{}.
	if s.AdditionalProperties.Has != nil && *s.AdditionalProperties.Has {
		return true
	}
	if s.AdditionalProperties.Schema != nil {
		return true
	}
	// additionalProperties 미지정 — OpenAPI 3.0 기본값은 허용.
	// oapi-codegen은 map[string]interface{}를 생성하므로 JSONB 취급.
	if s.AdditionalProperties.Has == nil && s.AdditionalProperties.Schema == nil {
		return true
	}
	return false
}
