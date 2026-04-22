//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what findPrecedingResource — path parts에서 param 앞 리소스명을 snake_case로 추출
package hurl

import "strings"

// findPrecedingResource walks back from idx to find the nearest literal
// path segment, strips a trailing "s" (simple plural), and returns it
// normalized to snake_case via snakeHurlName.
func findPrecedingResource(parts []string, idx int) string {
	for j := idx - 1; j >= 0; j-- {
		if parts[j] == "" || pathParamRe.MatchString(parts[j]) {
			continue
		}
		return snakeHurlName(strings.TrimSuffix(parts[j], "s"))
	}
	return ""
}
