//ff:func feature=gen-gogin type=util control=sequence
//ff:what isFloatHead — head 토큰이 float family 인지 판정

package types

// isFloatHead reports whether head is a float-family PG type token.
func isFloatHead(head string) bool {
	return floatHeads[head]
}
