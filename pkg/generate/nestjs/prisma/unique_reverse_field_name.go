//ff:func feature=gen-nestjs type=generator control=sequence
//ff:what uniqueReverseFieldName — 동일 타깃 다중 FK 시 중복 제거된 역관계 필드명 반환

package prisma

// uniqueReverseFieldName returns a deduplicated field name when multiple
// FKs point to the same target. Currently uses table name as-is since
// each table has a unique name.
func uniqueReverseFieldName(base string, _ int) string {
	return base
}
