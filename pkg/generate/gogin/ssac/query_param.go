//ff:type feature=gen-gogin type=model
//ff:what queryParam — OpenAPI query parameter 메타데이터 (enum alias, format, nullable)

package ssac

// queryParam describes one OpenAPI query parameter with metadata needed to
// generate correctly typed Go access code (enum alias types, int64 format,
// nullable markers). GoType is the normalized primitive name yongol uses
// internally: "integer", "integer64", "string", "bool", etc.
type queryParam struct {
	GoType       string // "integer" | "integer32" | "integer64" | "string" | "bool"
	IsEnum       bool
	EnumTypeName string // oapi-codegen alias type name, e.g. "ListAuditLogsParamsSortBy"
	IsRequired   bool
	IsNullable   bool
	Format       string // "int64" | "date-time" | ...
}
