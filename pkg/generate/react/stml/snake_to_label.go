//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what snake_case 문자열을 Title Case 라벨로 변환한다
package stml

import "strings"

func snakeToLabel(name string) string {
	parts := strings.Split(name, "_")
	for i, p := range parts {
		if len(p) > 0 {
			parts[i] = strings.ToUpper(p[:1]) + p[1:]
		}
	}
	return strings.Join(parts, " ")
}
