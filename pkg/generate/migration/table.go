//ff:type feature=migration type=model
//ff:what Table — 단일 CREATE TABLE 의 정규화된 AST 표현
package migration

// Table describes a single CREATE TABLE.
type Table struct {
	Name        string             // lowercase canonical name
	Columns     []*Column          // order preserved from DDL source
	PrimaryKey  []string           // columns forming the PK (canonical names)
	Indexes     []*Index           // UNIQUE inline + CREATE INDEX unified
	ForeignKeys []*ForeignKey      // inline REFERENCES + ALTER TABLE ADD FK unified
	Checks      []*CheckConstraint // CHECK (...) constraints
	Comment     string             // optional COMMENT ON TABLE (v1 WARN only)
	Sentinels   []SentinelInsert   // @sentinel INSERT blocks attached to this table
	errs        []string           // column-level parse errors (IDENTITY+DEFAULT conflict etc.)
}

// SentinelInsert captures one `-- @sentinel` INSERT statement verbatim,
// so the migration emitter can embed it unchanged after all CREATE TABLE
// and before CREATE INDEX / ALTER TABLE ADD FOREIGN KEY statements.
type SentinelInsert struct {
	SQL string // raw SQL through the final `;`
}
