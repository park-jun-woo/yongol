//ff:type feature=manifest type=model
//ff:what Table — DDL CREATE TABLE에서 추출한 테이블 메타데이터
package ddl

// Table holds parsed metadata for a single DDL table.
type Table struct {
	Name            string
	File            string              // 원본 .sql 파일 경로
	Line            int                 // CREATE TABLE 줄 번호 (1-based, 0 = 미상)
	Columns         map[string]string   // column_name → Go type
	ColumnOrder     []string            // DDL definition order
	ForeignKeys     []ForeignKey
	Indexes         []Index
	PrimaryKey      []string
	VarcharLen      map[string]int      // column → VARCHAR(N)
	CheckEnums      map[string][]string // column → CHECK IN values
	Archived         bool                // `-- @archived` directly above CREATE TABLE
	ArchivedColumns  map[string]bool     // columns marked with `-- @archived`
	SensitiveColumns map[string]bool     // columns marked with `-- @sensitive`
	Defaults         map[string]string   // column → DEFAULT '<value>' (string literal)
	NotNullCols      map[string]bool     // columns with explicit NOT NULL constraint
	NullableAnnot    map[string]bool     // columns marked with `-- @nullable`
}
