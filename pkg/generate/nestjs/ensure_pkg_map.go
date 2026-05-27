//ff:func feature=gen-nestjs type=util control=sequence
//ff:what ensurePkgMap — 패키지 맵 초기화 보장

package nestjs

// ensurePkgMap initializes a sub-map if it doesn't exist.
func ensurePkgMap(pm map[string]map[string]bool, key string) {
	if pm[key] == nil {
		pm[key] = make(map[string]bool)
	}
}
