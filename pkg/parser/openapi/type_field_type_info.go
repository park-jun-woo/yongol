//ff:type feature=openapi-parse type=model
//ff:what FieldTypeInfo — 응답 스키마 필드의 OpenAPI 타입과 포맷

package openapi

// FieldTypeInfo carries the OpenAPI type and format of a response schema
// field. The react emitter consults it for type-aware data-bind rendering:
// the Type drives the boolean/number branch and Format ("date"/"date-time")
// drives the locale-formatting branch (plans/gen/frontend Phase037, BUG-126).
type FieldTypeInfo struct {
	Type   string // OpenAPI type, e.g. "string", "integer", "boolean", "object", "array"
	Format string // OpenAPI format, e.g. "date", "date-time", "uri" ("" when absent)
}
