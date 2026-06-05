//ff:func feature=stml-parse type=parser control=selection dimension=1
//ff:what 연산자·논리·그룹 심볼 토큰을 읽는다 (&& || ! ( ) . = != > < >= <=)
package stml

import "fmt"

// lexGuardSymbol reads an operator/grouping symbol starting at index i. It
// returns the token and the index just past it, or an error for malformed
// symbols (e.g. a lone "&").
func lexGuardSymbol(runes []rune, i int) (guardToken, int, error) {
	r := runes[i]
	next := func(k int) rune {
		if i+k < len(runes) {
			return runes[i+k]
		}
		return 0
	}
	switch r {
	case '&':
		if next(1) == '&' {
			return guardToken{kind: tokAnd, text: "&&"}, i + 2, nil
		}
		return guardToken{}, 0, fmt.Errorf("invalid operator %q in guard expression (use &&)", "&")
	case '|':
		if next(1) == '|' {
			return guardToken{kind: tokOr, text: "||"}, i + 2, nil
		}
		return guardToken{}, 0, fmt.Errorf("invalid operator %q in guard expression (use ||)", "|")
	case '!':
		if next(1) == '=' {
			return guardToken{kind: tokOp, text: "!="}, i + 2, nil
		}
		return guardToken{kind: tokNot, text: "!"}, i + 1, nil
	case '(':
		return guardToken{kind: tokLParen, text: "("}, i + 1, nil
	case ')':
		return guardToken{kind: tokRParen, text: ")"}, i + 1, nil
	case '.':
		return guardToken{kind: tokDot, text: "."}, i + 1, nil
	case '=':
		return guardToken{kind: tokOp, text: "="}, i + 1, nil
	case '>', '<':
		if next(1) == '=' {
			return guardToken{kind: tokOp, text: string(r) + "="}, i + 2, nil
		}
		return guardToken{kind: tokOp, text: string(r)}, i + 1, nil
	default:
		return guardToken{}, 0, fmt.Errorf("unexpected symbol %q in guard expression", string(r))
	}
}
