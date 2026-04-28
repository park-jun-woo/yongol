//ff:type feature=manifest type=model
//ff:what Table — table metadata extracted from a DDL CREATE TABLE statement
package ddl

// Table holds parsed metadata for a single DDL table. Per-column metadata
// (NotNull, DEFAULT, VARCHAR(N), CHECK IN, @archived, @sensitive,
// @nullable) is consolidated inside Column; see column.go.
type Table struct {
	Name        string
	File        string            // path of the source .sql file
	Line        int               // CREATE TABLE line number (1-based, 0 = unknown)
	Columns     map[string]Column // column_name → Column metadata
	ColumnOrder []string          // DDL definition order
	ForeignKeys []ForeignKey
	Indexes     []Index
	PrimaryKey  []string
	Archived    bool             // `-- @archived` directly above CREATE TABLE
	Sentinels   []SentinelInsert // `-- @sentinel` INSERT blocks targeting this table
}
