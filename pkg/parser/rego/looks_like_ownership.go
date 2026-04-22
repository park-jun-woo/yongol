//ff:func feature=policy type=util control=sequence
//ff:what looksLikeOwnership — 주석이 @ownership 을 의도했는지 판별

package rego

import "strings"

// looksLikeOwnership returns true if the trimmed comment line starts with
// "#<space>*@ownership<space>" (or is just "@ownership") but fails the strict
// reOwnership regex. Used to surface malformed annotations as parse ERRORs
// instead of silently dropping them.
func looksLikeOwnership(line string) bool {
	if !strings.HasPrefix(line, "#") {
		return false
	}
	body := strings.TrimLeft(strings.TrimPrefix(line, "#"), " \t")
	return body == "@ownership" || strings.HasPrefix(body, "@ownership ")
}
