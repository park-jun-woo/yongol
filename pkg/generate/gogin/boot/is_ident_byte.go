//ff:func feature=gen-gogin type=util control=selection
//ff:what isIdentByte — ASCII letter/digit/underscore 판정 (import usage 단어경계)

package boot

// isIdentByte reports whether b could be part of a Go identifier —
// letters, digits, and underscore. Used to enforce the left-side word
// boundary for import usage detection. Only ASCII is considered because
// the generator emits ASCII identifiers only; if the generator ever
// produces unicode identifiers this needs to widen.
func isIdentByte(b byte) bool {
	switch {
	case b >= 'a' && b <= 'z':
		return true
	case b >= 'A' && b <= 'Z':
		return true
	case b >= '0' && b <= '9':
		return true
	case b == '_':
		return true
	}
	return false
}
