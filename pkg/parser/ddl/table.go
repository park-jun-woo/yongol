//ff:type feature=manifest type=model
//ff:what Table — table metadata extracted from a DDL CREATE TABLE statement
package ddl

// Table holds parsed metadata for a single DDL table.
type Table struct {
	Name             string
	File             string            // path of the source .sql file
	Line             int               // CREATE TABLE line number (1-based, 0 = unknown)
	Columns          map[string]string // column_name → Go type
	ColumnOrder      []string          // DDL definition order
	ForeignKeys      []ForeignKey
	Indexes          []Index
	PrimaryKey       []string
	VarcharLen       map[string]int      // column → VARCHAR(N)
	CheckEnums       map[string][]string // column → CHECK IN values
	Archived         bool                // `-- @archived` directly above CREATE TABLE
	ArchivedColumns  map[string]bool     // columns marked with `-- @archived`
	SensitiveColumns map[string]bool     // columns marked with `-- @sensitive`
	Defaults         map[string]string   // column → DEFAULT '<value>' (string literal)
	NotNullCols      map[string]bool     // columns with explicit NOT NULL constraint
	NullableAnnot    map[string]bool     // columns marked with `-- @nullable`
	Sentinels        []SentinelInsert    // `-- @sentinel` INSERT blocks targeting this table
}
