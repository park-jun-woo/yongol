//ff:type feature=gen-ir type=model
//ff:what TypeBinding — 백엔드 언어 독립적 PG 컬럼 타입 바인딩 (DB/API 레이어 타입 + 변환 표현식)

package ir

import "github.com/park-jun-woo/yongol/pkg/generate/typemap"

// TypeBinding holds the language-agnostic type mapping for a single DDL
// column across all backend targets. Each backend's TypeRegistry
// implementation produces a TypeBinding from a PGFamily + BindOpts pair.
//
// Template placeholders follow the same grammar as GoTypeBinding:
//
//	{row}   → Row variable name (DB row → API model convert)
//	{field} → field name (PascalCase)
//	{var}   → source variable name (INSERT / response emit)
type TypeBinding struct {
	// Family is the PGFamily classification of the source column.
	Family typemap.PGFamily

	// NotNull reflects whether the column is effectively NOT NULL.
	NotNull bool

	// DB layer type (ORM / query builder target).
	DBType    string   // Go: "int64", TS: "number", Py: "int"
	DBImports []string // import paths for DBType usage

	// API layer type (request / response surface).
	APIType    string   // Go: "*openapi_types.UUID", TS: "string", Py: "UUID"
	APIImports []string // import paths for APIType usage

	// Conversion expressions between DB ↔ API layers.
	ToDBExpr       string // value → DB insert; placeholder: {var}
	ToAPIExpr      string // DB row → API model; placeholder: {row}, {field}
	ToResponseExpr string // model → JSON response; placeholder: {var}, {field}

	// Nil guard + validation.
	NilCheckExpr string // @empty/@exists nil guard predicate
	Supported    bool   // false → validate D-11 rejects this type

	// DefaultLiteral is the emit-time literal for NOT NULL DEFAULT columns.
	DefaultLiteral string
}
