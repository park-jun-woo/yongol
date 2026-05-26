//ff:func feature=gen-ir type=generator control=sequence
//ff:what buildRateLimitConfig -- manifest.backend.rate_limit → RateLimitConfig 변환

package ir

import "github.com/park-jun-woo/yongol/pkg/yongol"

// buildRateLimitConfig returns nil when no rate limit rules are defined.
func buildRateLimitConfig(fs *yongol.Fullstack) *RateLimitConfig {
	if fs == nil || fs.Manifest == nil || len(fs.Manifest.Backend.RateLimit) == 0 {
		return nil
	}
	// We record the presence of rate-limit config; detailed route
	// resolution requires OpenAPI data and belongs in the renderer.
	cfg := &RateLimitConfig{}
	return cfg
}
