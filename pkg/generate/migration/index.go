//ff:type feature=migration type=model
//ff:what Index — 인덱스(고유/일반) 정의 AST
package migration

// Index describes a unique or non-unique btree index.
type Index struct {
	Name    string
	Columns []string
	Unique  bool
	Where   string // partial index predicate (canonical), "" when absent
}
