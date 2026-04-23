//ff:type feature=migration type=model topic=migration-hints
//ff:what Hints — DDL 주석 힌트 전체 모음 (rename/cast/backfill/data_migration/allow_destructive)
package migration

// Hints carries DDL comment hints consumed by Diff / safety checks.
type Hints struct {
	RenameTables     []RenameTableHint
	RenameColumns    []RenameColumnHint
	Casts            map[colKey]string // (table,col) → USING expr
	Backfills        map[colKey]string // (table,col) → literal default
	DataMigrations   map[string]string // table → sidecar file path
	AllowDestructive map[string]bool   // table → allow DROP TABLE/COLUMN
}
