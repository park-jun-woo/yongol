//ff:type feature=manifest type=model
//ff:what Index — DDL 인덱스 정보
package ddl

// Index represents a database index.
type Index struct {
	Name     string
	Columns  []string
	IsUnique bool
	// Method stores the optional `USING <method>` clause (e.g. "btree",
	// "gin", "gist", "brin", "hash", "spgist"). Empty string when the
	// DDL omits the clause.
	Method string
}
