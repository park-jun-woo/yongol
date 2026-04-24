//ff:type feature=migration type=model
//ff:what Index — 인덱스(고유/일반) 정의 AST
package migration

// Index describes a unique or non-unique index. Method holds the optional
// `USING <method>` clause (e.g. "gin", "gist", "brin", "hash", "spgist").
// Empty string or "btree" are treated equivalently when emitted — PostgreSQL
// default is btree, so the emitter omits `USING btree` for brevity.
type Index struct {
	Name    string
	Columns []string
	Unique  bool
	Method  string // "", "btree", "gin", "gist", "brin", "hash", "spgist"
	Where   string // partial index predicate (canonical), "" when absent
}
