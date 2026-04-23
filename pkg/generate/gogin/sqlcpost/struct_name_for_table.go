//ff:func feature=gen-gogin type=util control=sequence
//ff:what StructNameForTable — DDL 테이블명 → sqlc Go struct 이름 (exported wrapper)

package sqlcpost

// StructNameForTable is the exported form of structNameFor so other
// generate sub-packages (e.g. ssac for converter emission) can reason
// about which sqlc row types actually exist at build time. Consolidating
// this mapping in one place keeps the singular/PascalCase rules in sync
// between the sqlc post-processing pass and the converter emitter.
func StructNameForTable(table string) string {
	return structNameFor(table)
}
