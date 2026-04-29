//ff:func feature=gen-gogin type=util control=sequence
//ff:what isStringHead — head 토큰이 string family 인지 판정

package types

// isStringHead reports whether head is a string-family PG type token.
func isStringHead(head string) bool {
	return stringHeads[head]
}
