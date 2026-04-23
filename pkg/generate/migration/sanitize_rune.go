//ff:func feature=migration type=util control=selection
//ff:what sanitizeRune — 한 룬을 sanitize 규칙에 따라 치환 (a-z0-9_ 유지, 외엔 _)
package migration

// sanitizeRune returns r itself when it's an allowed char ([a-z0-9_]),
// otherwise '_'.
func sanitizeRune(r rune) rune {
	switch {
	case r >= 'a' && r <= 'z':
		return r
	case r >= '0' && r <= '9':
		return r
	case r == '_':
		return '_'
	}
	return '_'
}
