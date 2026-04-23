//ff:type feature=migration type=model
//ff:what ForeignKey — FK 제약 정의 (inline REFERENCES + ALTER TABLE ADD FK 통합)
package migration

// ForeignKey describes a FK constraint, either inline or from ALTER TABLE.
type ForeignKey struct {
	Name       string // PostgreSQL auto name when user omitted it
	Columns    []string
	RefTable   string
	RefColumns []string
	OnDelete   string // "CASCADE" | "SET NULL" | "RESTRICT" | "NO ACTION" — "" means default NO ACTION
	OnUpdate   string
}
