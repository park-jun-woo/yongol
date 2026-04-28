//ff:func feature=migration type=util control=sequence
//ff:what normalizeIndexMethod — "" 와 "btree" 를 동일 토큰으로 canonicalise

package migration

// normalizeIndexMethod canonicalises empty string and "btree" to the same
// token so equality treats them interchangeably.
func normalizeIndexMethod(m string) string {
	if m == "" {
		return "btree"
	}
	return m
}
