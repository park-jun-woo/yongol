//ff:func feature=validate type=util control=sequence dimension=1 topic=openapi-structural
//ff:what hasJSONContentWithSchema — response 가 application/json + schema 를 보유했는지 확인

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// hasJSONContentWithSchema returns true when the given response declares
// `content: application/json` with a non-nil schema. Used by O-5 to enforce
// that every 4xx/5xx response carries a structured JSON body so that
// oapi-codegen emits a `<Op><Status>JSONResponse` type (matching yongol's
// SSaC handler call sites).
//
// Returns false when the ResponseRef itself is nil, when its Value is nil,
// when no Content is declared, when the application/json media type is
// missing, or when the schema reference is empty.
func hasJSONContentWithSchema(ref *openapi3.ResponseRef) bool {
	if ref == nil || ref.Value == nil {
		return false
	}
	if ref.Value.Content == nil {
		return false
	}
	mt := ref.Value.Content.Get("application/json")
	if mt == nil || mt.Schema == nil {
		return false
	}
	// kin-openapi exposes both inline (Value) and reference (Ref) schemas.
	// Either form satisfies the rule — yongol does not require an inline
	// definition; a $ref to components.schemas.Error is the recommended
	// form.
	if mt.Schema.Value == nil && mt.Schema.Ref == "" {
		return false
	}
	return true
}
