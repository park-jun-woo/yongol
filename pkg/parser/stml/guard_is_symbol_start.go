//ff:func feature=stml-parse type=util control=sequence dimension=1
//ff:what 토큰 시작 문자가 연산자/그룹 심볼인지 판별한다
package stml

import "strings"

// isGuardSymbolStart reports whether r can begin an operator/grouping symbol.
func isGuardSymbolStart(r rune) bool {
	return strings.ContainsRune("&|!()=<>.", r)
}
