//ff:type feature=migration type=model
//ff:what Schema — diff 비교용 정규화된 PostgreSQL 스키마 AST 루트
package migration

// Schema is the canonical in-memory representation of a whole DDL
// directory. diff engine consumes two Schema values (prev, curr).
type Schema struct {
	// Tables keyed by lowercase canonical table name.
	Tables map[string]*Table
}
