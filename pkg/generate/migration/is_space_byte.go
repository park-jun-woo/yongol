//ff:func feature=migration type=util control=sequence
//ff:what isSpaceByte — SQL 토크나이저용 공백 문자 판정 (' '/'\t'/'\n'/'\r')
package migration

// isSpaceByte reports whether c is SQL whitespace.
func isSpaceByte(c byte) bool {
	return c == ' ' || c == '\t' || c == '\n' || c == '\r'
}
