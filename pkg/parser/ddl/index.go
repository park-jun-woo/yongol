//ff:type feature=manifest type=model
//ff:what Index — DDL 인덱스 정보
package ddl

// Index represents a database index.
type Index struct {
	Name     string
	Columns  []string
	IsUnique bool
}
