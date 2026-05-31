//ff:func feature=validate type=util control=sequence topic=func-check
//ff:what isUUIDType — report whether a Go type name is the oapi-codegen UUID type (excludes pgtype.UUID)

package ssac_func

// isUUIDType reports whether t names the oapi-codegen runtime UUID type. Both
// "openapi_types.UUID" and "types.UUID" refer to the same
// github.com/oapi-codegen/runtime/types.UUID under different import aliases.
// "pgtype.UUID" (the DB/sqlc type) is intentionally NOT matched here.
func isUUIDType(t string) bool {
	return t == "openapi_types.UUID" || t == "types.UUID"
}
