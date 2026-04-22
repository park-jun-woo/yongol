//ff:func feature=gen-filefunc type=util control=sequence
//ff:what insertFeatureIfNew — 비어있지 않은 이름이고 기존 키에 없을 때만 맵에 삽입
package filefunc

// insertFeatureIfNew adds (name, desc) to dst when name is non-empty and
// not already present as a key. Descriptions never overwrite existing
// entries.
func insertFeatureIfNew(dst map[string]string, name, desc string) {
	if name == "" {
		return
	}
	if _, ok := dst[name]; ok {
		return
	}
	dst[name] = desc
}
