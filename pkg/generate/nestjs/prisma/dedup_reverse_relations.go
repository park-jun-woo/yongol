//ff:func feature=gen-nestjs type=generator control=iteration dimension=1
//ff:what dedupReverseRelations — 동일 source 모델 중복 역관계 엔트리 제거

package prisma

// dedupReverseRelations removes duplicate entries (same source model).
func dedupReverseRelations(rels []reverseRelation) []reverseRelation {
	seen := make(map[string]bool)
	var result []reverseRelation
	for _, r := range rels {
		key := r.FieldName + ":" + r.ModelName
		if seen[key] {
			continue
		}
		seen[key] = true
		result = append(result, r)
	}
	return result
}
