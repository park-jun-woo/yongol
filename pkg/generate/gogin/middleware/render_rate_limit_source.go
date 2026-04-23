//ff:func feature=gen-gogin type=util control=sequence
//ff:what renderRateLimitSource — 템플릿의 __MODULE__ 토큰을 실제 모듈 경로로 치환

package middleware

import "strings"

// renderRateLimitSource substitutes __MODULE__ with the actual module path.
func renderRateLimitSource(modulePath string) string {
	return strings.ReplaceAll(rateLimitSourceTemplate, "__MODULE__", modulePath)
}
