//ff:type feature=validate type=model topic=openapi-ddl
//ff:what openAPITypeCompat — DDL Go 타입에 대응하는 기대 OpenAPI type/format

package openapi_ddl

// openAPITypeCompat holds the expected OpenAPI type/format for a given DDL Go type.
type openAPITypeCompat struct {
	oType   string // expected OpenAPI type (e.g. "integer", "string", "boolean", "number")
	oFormat string // expected OpenAPI format (empty = any format accepted; non-empty = must match)
}
