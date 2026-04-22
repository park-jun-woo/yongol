//ff:func feature=gen-filefunc type=util control=sequence
//ff:what upsertFeatureDesc — 기존 설명이 비어있을 때만 새 설명으로 덮어쓴다
package filefunc

// upsertFeatureDesc inserts (name, desc) into dst when name is non-empty.
// If the key already exists with a non-empty description, the existing
// description is kept; otherwise desc (which may itself be empty) wins.
func upsertFeatureDesc(dst map[string]string, name, desc string) {
	if name == "" {
		return
	}
	cur, ok := dst[name]
	if ok && cur != "" {
		return
	}
	dst[name] = desc
}
