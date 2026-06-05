//ff:func feature=stml-parse type=parser control=selection dimension=1
//ff:what 인덱스 i에서 토큰 하나를 분류·추출한다 (공백은 emit=false로 건너뜀)
package stml

import (
	"fmt"
	"unicode"
)

// lexGuardStep classifies and reads a single token at index i. The emit return
// is false for skipped whitespace; next is the index to continue from.
func lexGuardStep(runes []rune, i int) (tok guardToken, emit bool, next int, err error) {
	r := runes[i]
	switch {
	case unicode.IsSpace(r):
		return guardToken{}, false, i + 1, nil
	case r == '\'' || r == '"':
		tok, next, err = lexGuardString(runes, i)
		return tok, err == nil, next, err
	case isGuardSymbolStart(r):
		tok, next, err = lexGuardSymbol(runes, i)
		return tok, err == nil, next, err
	case isGuardIdentRune(r):
		tok, next = lexGuardIdent(runes, i)
		return tok, true, next, nil
	default:
		return guardToken{}, false, i, fmt.Errorf("unexpected character %q in guard expression", string(r))
	}
}
