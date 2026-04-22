//ff:type feature=manifest type=model
//ff:what ForeignKey — 외래키 관계
package ddl

// ForeignKey represents a foreign key relationship.
type ForeignKey struct {
	Column    string
	RefTable  string
	RefColumn string
}
