//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what 따옴표 문자열 리터럴 토큰을 읽는다 (' 또는 " 로 둘러싸인 값)
package stml

import "fmt"

// lexGuardString reads a quoted string literal starting at index i (a quote
// rune). It returns the token, the index just past the closing quote, or an
// error if the string is unterminated.
func lexGuardString(runes []rune, i int) (guardToken, int, error) {
	quote := runes[i]
	j := i + 1
	for j < len(runes) {
		if runes[j] == quote {
			text := string(runes[i+1 : j])
			return guardToken{kind: tokString, text: text}, j + 1, nil
		}
		j++
	}
	return guardToken{}, 0, fmt.Errorf("unterminated string literal in guard expression")
}
