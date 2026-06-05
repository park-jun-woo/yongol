//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what 식별자/숫자/enum-literal 토큰을 읽는다 (영숫자·_·- 연속)
package stml

// lexGuardIdent reads an identifier/number/enum-literal token starting at index
// i. It returns the token and the index just past the last consumed rune.
func lexGuardIdent(runes []rune, i int) (guardToken, int) {
	j := i
	for j < len(runes) && isGuardIdentRune(runes[j]) {
		j++
	}
	return guardToken{kind: tokIdent, text: string(runes[i:j])}, j
}
