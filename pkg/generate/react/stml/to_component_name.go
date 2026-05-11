//ff:func feature=stml-gen type=util control=iteration dimension=1 topic=string-convert
//ff:what 케밥케이스 이름을 PascalCase 컴포넌트 이름으로 변환한다
package stml

import "strings"

// toComponentName converts "my-reservations-page" to "MyReservationsPage".
func toComponentName(name string) string {
	parts := strings.Split(name, "-")
	for i, p := range parts {
		parts[i] = toUpperFirst(p)
	}
	return strings.Join(parts, "")
}
