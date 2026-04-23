//ff:type feature=migration type=model
//ff:what CanonicalType — 타입 정규화 (int4/INTEGER/integer → INTEGER)
package migration

// CanonicalType represents a PostgreSQL column type in a canonical form
// so that "int4", "INTEGER", "integer" all compare equal.
type CanonicalType struct {
	Base      string // "INTEGER" / "VARCHAR" / "TEXT" / "BOOLEAN" / "TIMESTAMPTZ" / ...
	Length    int    // VARCHAR(N), CHAR(N) — 0 when not applicable
	Precision int    // NUMERIC(p,s) — 0 when not applicable
	Scale     int    // NUMERIC(p,s) — 0 when not applicable
	Array     bool   // true for INTEGER[] etc
}
