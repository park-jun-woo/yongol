//ff:func feature=validate type=util control=sequence topic=stml-design
//ff:what isOverrideComment — 주석 데이터가 @override로 시작하는지 확인
package stml_design

import (
	"strings"
)

// isOverrideComment checks if a comment node's data starts with "@override".
func isOverrideComment(data string) bool {
	return strings.HasPrefix(strings.TrimSpace(data), "@override")
}
