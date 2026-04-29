//ff:func feature=gen-gogin type=util control=sequence
//ff:what isBooleanHead — head 토큰이 boolean family 인지 판정

package types

// isBooleanHead reports whether head is a boolean-family PG type token.
func isBooleanHead(head string) bool {
	return booleanHeads[head]
}
