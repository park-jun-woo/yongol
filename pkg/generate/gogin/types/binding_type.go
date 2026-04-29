//ff:type feature=gen-gogin type=model
//ff:what GoTypeBinding — PG 컬럼 → Go 타입 매핑 결과 (sqlc / api / convert / insert / response 표현 통합)

package types

// GoTypeBinding holds every Go-side decision derived from a single parsed
// DDL column: the sqlc Go type, the api struct field type, import wiring,
// and three template strings for the convert / insert / response emit
// sites. The struct is the single source of truth shared between codegen
// (pkg/generate/gogin/ssac) and validate (Q-12 + per-type Q-NN family).
//
// Template grammar (see Expand):
//
//	{row}   → Row variable name (Row → Model convert)
//	{field} → DB / Model field name (PascalCase; sqlc convention)
//	{var}   → source / model variable name (INSERT / response)
//
// A literal `{` is escaped as `{{`. Unused placeholders are silently
// substituted with empty string by Expand so emit sites do not need to
// branch on Kind for placeholder presence.
type GoTypeBinding struct {
	// SqlcGoType is the canonical Go type sqlc pgx/v5 should emit for this
	// column. For NeedsOverride=true types this is the value that must
	// appear in sqlc.yaml's `go_type` override (e.g. "pgtype.UUID"). For
	// natively supported types it is the primitive name (e.g. "int64").
	SqlcGoType string

	// NeedsOverride is true when sqlc.yaml MUST register an explicit
	// pgtype override for this column. validate Q-12 (UUID) and the
	// per-type Q-NN family share this flag.
	NeedsOverride bool

	// ApiField is the api struct field type (oapi-codegen surface),
	// e.g. "string", "*string", "openapi_types.UUID", "[]string".
	ApiField string

	// Imports lists fully-qualified import paths the emit sites must wire
	// in for ConvertExpr / InsertExpr / ResponseExpr to compile.
	Imports []string

	// ConvertExpr is the Row → Model template. The convert emitter calls
	// Expand(ConvertExpr, row, field, "").
	ConvertExpr string

	// InsertExpr is the value → sqlc params template. The INSERT emitter
	// calls Expand(InsertExpr, "", field, varName) where varName is the
	// source expression (request body field, default literal wrap, …).
	InsertExpr string

	// ResponseExpr is the Model → JSON template. The response emitter
	// calls Expand(ResponseExpr, "", field, modelVar) to extract the
	// primitive payload from any pgtype wrapper.
	ResponseExpr string

	// DefaultLiteral is the emit-time literal used for NOT NULL DEFAULT
	// columns when no caller value is supplied. Empty string means the
	// caller must always provide a value.
	DefaultLiteral string

	// Kind is the binding family (used for branching that cannot be
	// expressed via templates).
	Kind BindingKind

	// Supported is false only for KindUnsupported. validate D-11 emits
	// an ERROR per such column before generate runs.
	Supported bool
}
