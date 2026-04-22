//ff:func feature=manifest type=util control=iteration dimension=1
//ff:what findTableKeyword — parts에서 TABLE 키워드의 인덱스 반환
package ddl

import "strings"

func findTableKeyword(parts []string) int {
	for i, p := range parts {
		if strings.ToUpper(p) == "TABLE" {
			return i
		}
	}
	return -1
}
