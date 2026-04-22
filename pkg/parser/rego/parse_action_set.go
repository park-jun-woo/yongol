//ff:func feature=policy type=util control=iteration dimension=1
//ff:what parseActionSet — 쉼표 구분된 액션 문자열을 슬라이스로 파싱
package rego

import "strings"

func parseActionSet(s string) []string {
	var actions []string
	for _, part := range strings.Split(s, ",") {
		part = strings.TrimSpace(part)
		part = strings.Trim(part, "\"")
		if part != "" {
			actions = append(actions, part)
		}
	}
	return actions
}
