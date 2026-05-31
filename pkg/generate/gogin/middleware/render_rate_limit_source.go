//ff:func feature=gen-gogin type=util control=sequence topic=rate-limit
//ff:what renderRateLimitSources — rate_limit 템플릿을 파일별 소스 맵으로 생성

package middleware

// renderRateLimitSources returns a map of filename to source content for
// the rate-limit middleware split into 1-file-1-func.
func renderRateLimitSources() map[string]string {
	return map[string]string{
		"fixed_rate_limit.go":     rateLimitFixedSource,
		"fixed_rate_limit_key.go": rateLimitKeySource,
		"route_rate_limit.go":     routeRateLimitSource,
		"route_rate_limit_key.go": routeRateLimitKeySource,
	}
}
