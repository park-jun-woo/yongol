//ff:type feature=gen-ir type=model
//ff:what BodyFieldMeta -- 단일 OpenAPI 요청 본문 property 메타데이터

package ir

// BodyFieldMeta holds metadata for a single OpenAPI request body property.
type BodyFieldMeta struct {
	// Name is the OpenAPI property name.
	Name string
	// Required is true when the property is listed in the schema's required array.
	Required bool
	// Format is the OpenAPI format (e.g. "email", "uuid", "date-time", "enum", "").
	Format string
}
