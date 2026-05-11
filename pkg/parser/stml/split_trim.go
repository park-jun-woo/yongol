//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what 쉼표 구분 문자열을 분리하고 공백 제거
package stml

import "strings"

// splitTrim splits a comma-separated string and trims whitespace.
func splitTrim(v string) []string {
	raw := strings.Split(v, ",")
	var result []string
	for _, s := range raw {
		s = strings.TrimSpace(s)
		if s != "" {
			result = append(result, s)
		}
	}
	return result
}
