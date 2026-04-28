//ff:type feature=manifest type=model
//ff:what Column — DDL 컬럼 메타데이터 통합 구조체 (RawType 보존 + 8 평행 맵 흡수)
package ddl

// Column holds the parsed metadata for a single column definition. The raw
// PostgreSQL type token is preserved verbatim (e.g. "BIGINT", "TEXT[]",
// "VARCHAR(255)", "NUMERIC(10,2)") so downstream consumers can derive Go
// type projections, array marker decisions, and length constraints from a
// single source of truth.
//
// This struct replaces the eight parallel maps that previously decorated
// Table (NotNullCols, NullableAnnot, VarcharLen, CheckEnums, Defaults,
// ArchivedColumns, SensitiveColumns plus the lossy Columns map[string]string).
type Column struct {
	Name           string
	RawType        string   // verbatim PG type token, e.g. "BIGINT", "VARCHAR(255)", "TEXT[]"
	NotNull        bool     // explicit NOT NULL or PRIMARY KEY
	NullableAnnot  bool     // `-- @nullable` annotation
	HasDefault     bool     // DEFAULT '<literal>' present (string-literal only)
	DefaultLiteral string   // captured literal value
	VarcharLen     int      // 0 if not VARCHAR(N)
	CheckEnum      []string // empty if no CHECK IN (...)
	Archived       bool     // `-- @archived`
	Sensitive      bool     // `-- @sensitive`
}
