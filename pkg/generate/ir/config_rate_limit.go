//ff:type feature=gen-ir type=model
//ff:what RateLimitConfig -- 라우트별 rate limit 설정

package ir

// RateLimitConfig holds the resolved rate-limiting rules.
type RateLimitConfig struct {
	// Rules maps route paths to their rate limit configuration.
	Rules []RateLimitRule
}
