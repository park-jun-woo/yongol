//ff:func feature=gen-gogin type=util control=selection
//ff:what apiCastFor — sqlc string 컬럼 → api-side named type 캐스트 표현식 선택 (enum/format)

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// apiCastFor returns the Go expression (including package prefix) to
// cast a sqlc string column into the api-side named type, or "" when no
// cast is required. Two families are supported:
//
//   - String enums → `api.<Parent><Property>` (oapi-codegen's default
//     naming, e.g. WorkflowStatus for Workflow.status with enum values)
//   - String formats that oapi-codegen maps to a runtime wrapper type —
//     format: email → `openapi_types.Email` and format: uuid →
//     `openapi_types.UUID`. The runtime package is
//     github.com/oapi-codegen/runtime/types, already imported by
//     generated api files; convert files gain the import only when
//     at least one property requires it.
//
// Returns "" when the schema is missing, not a string, or has neither
// an enum nor a recognised format.
func apiCastFor(parentName, jsonName string, ref *openapi3.SchemaRef) string {
	if ref == nil || ref.Value == nil {
		return ""
	}
	s := ref.Value
	if s.Type != nil && !s.Type.Is("string") {
		return ""
	}
	if len(s.Enum) > 0 {
		return "api." + parentName + pascalCase(jsonName)
	}
	switch s.Format {
	case "email":
		return "openapi_types.Email"
	case "uuid":
		return "openapi_types.UUID"
	}
	return ""
}
