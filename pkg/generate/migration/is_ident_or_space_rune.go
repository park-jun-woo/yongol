//ff:func feature=migration type=util control=selection
//ff:what isIdentOrSpaceRune — 식별자 룬 ([A-Za-z0-9_]) 또는 공백 판정 ("character varying" 대응)
package migration

// isIdentOrSpaceRune reports whether r can appear inside a cast-target
// identifier like "character varying".
func isIdentOrSpaceRune(r rune) bool {
	switch {
	case r == ' ' || r == '_':
		return true
	case r >= 'a' && r <= 'z':
		return true
	case r >= 'A' && r <= 'Z':
		return true
	case r >= '0' && r <= '9':
		return true
	}
	return false
}
