//ff:func feature=gen-gogin type=util control=sequence
//ff:what isIntegerHead — head 토큰이 integer family 인지 판정

package types

// isIntegerHead reports whether head is an integer-family PG type token.
func isIntegerHead(head string) bool {
	return integerHeads[head]
}
