//ff:type feature=validate type=model topic=stml-openapi
//ff:what responseFieldInfo — 응답 스키마 필드의 이름·타입 정보

package stml_openapi

// responseFieldInfo holds the type metadata for a response schema field.
type responseFieldInfo struct {
	typ string // "string", "integer", "array", "object", etc.
}
