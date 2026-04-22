//ff:func feature=validate type=util control=sequence topic=openapi-structural
//ff:what isPasswordField — 필드명이 password 패턴인지 확인

package openapi

import "strings"

func isPasswordField(name string) bool {
	lower := strings.ToLower(name)
	return lower == "password" || strings.HasSuffix(lower, "password") || strings.HasPrefix(lower, "password")
}
