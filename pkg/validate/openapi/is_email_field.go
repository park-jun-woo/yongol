//ff:func feature=validate type=util control=sequence topic=openapi-structural
//ff:what isEmailField — 필드명이 email 패턴인지 확인

package openapi

import "strings"

func isEmailField(name string) bool {
	lower := strings.ToLower(name)
	return lower == "email" || strings.HasSuffix(lower, "email") || strings.HasPrefix(lower, "email")
}
