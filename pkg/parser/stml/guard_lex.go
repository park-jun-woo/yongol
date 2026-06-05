//ff:func feature=stml-parse type=parser control=iteration dimension=1
//ff:what 가드 조건 문자열을 토큰 슬라이스로 분해한다 (EBNF 어휘 토큰)
package stml

// lexGuard splits a guard condition string into tokens. It returns an error for
// any character outside the guard lexicon (§3.4).
func lexGuard(s string) ([]guardToken, error) {
	var toks []guardToken
	runes := []rune(s)
	i := 0
	for i < len(runes) {
		tok, emit, next, err := lexGuardStep(runes, i)
		if err != nil {
			return nil, err
		}
		if emit {
			toks = append(toks, tok)
		}
		i = next
	}
	toks = append(toks, guardToken{kind: tokEOF})
	return toks, nil
}
