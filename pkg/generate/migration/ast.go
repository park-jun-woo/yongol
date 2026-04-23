//ff:func feature=migration type=model control=sequence
//ff:what Schema AST — diff 비교용 정규화된 PostgreSQL 스키마 표현
package migration

// Schema is the canonical in-memory representation of a whole DDL
// directory. diff engine consumes two Schema values (prev, curr).
type Schema struct {
	// Tables keyed by lowercase canonical table name.
	Tables map[string]*Table
}

// Table describes a single CREATE TABLE.
type Table struct {
	Name        string             // lowercase canonical name
	Columns     []*Column          // order preserved from DDL source
	PrimaryKey  []string           // columns forming the PK (canonical names)
	Indexes     []*Index           // UNIQUE inline + CREATE INDEX unified
	ForeignKeys []*ForeignKey      // inline REFERENCES + ALTER TABLE ADD FK unified
	Checks      []*CheckConstraint // CHECK (...) constraints
	Comment     string             // optional COMMENT ON TABLE (v1 WARN only)
}

// Column is a single column definition, already normalised.
type Column struct {
	Name     string
	Type     CanonicalType
	Nullable bool
	Default  string // canonical default expression, "" = none
	Comment  string // trailing -- comment (@sensitive / @nullable are consumed separately)
}

// CanonicalType represents a PostgreSQL column type in a canonical form
// so that "int4", "INTEGER", "integer" all compare equal.
type CanonicalType struct {
	Base      string // "INTEGER" / "VARCHAR" / "TEXT" / "BOOLEAN" / "TIMESTAMPTZ" / ...
	Length    int    // VARCHAR(N), CHAR(N) — 0 when not applicable
	Precision int    // NUMERIC(p,s) — 0 when not applicable
	Scale     int    // NUMERIC(p,s) — 0 when not applicable
	Array     bool   // true for INTEGER[] etc
}

// Index describes a unique or non-unique btree index.
type Index struct {
	Name    string
	Columns []string
	Unique  bool
	Where   string // partial index predicate (canonical), "" when absent
}

// ForeignKey describes a FK constraint, either inline or from ALTER TABLE.
type ForeignKey struct {
	Name       string // PostgreSQL auto name when user omitted it
	Columns    []string
	RefTable   string
	RefColumns []string
	OnDelete   string // "CASCADE" | "SET NULL" | "RESTRICT" | "NO ACTION" — "" means default NO ACTION
	OnUpdate   string
}

// CheckConstraint describes a CHECK (...) constraint.
type CheckConstraint struct {
	Name       string // auto-generated when user omitted
	Expression string // canonical expression text
}

// Equal reports whether two CanonicalType values are identical in all
// fields. Helper for diff engines (Phase003).
func (t CanonicalType) Equal(other CanonicalType) bool {
	return t.Base == other.Base &&
		t.Length == other.Length &&
		t.Precision == other.Precision &&
		t.Scale == other.Scale &&
		t.Array == other.Array
}

// SQL renders a CanonicalType back to PostgreSQL DDL fragment.
func (t CanonicalType) SQL() string {
	s := t.Base
	switch {
	case t.Length > 0 && t.Base == "VARCHAR":
		s = "VARCHAR(" + itoa(t.Length) + ")"
	case t.Length > 0 && t.Base == "CHAR":
		s = "CHAR(" + itoa(t.Length) + ")"
	case t.Precision > 0 && t.Base == "NUMERIC":
		if t.Scale > 0 {
			s = "NUMERIC(" + itoa(t.Precision) + "," + itoa(t.Scale) + ")"
		} else {
			s = "NUMERIC(" + itoa(t.Precision) + ")"
		}
	}
	if t.Array {
		s += "[]"
	}
	return s
}

// NewSchema returns an empty Schema with a non-nil Tables map.
func NewSchema() *Schema {
	return &Schema{Tables: map[string]*Table{}}
}
