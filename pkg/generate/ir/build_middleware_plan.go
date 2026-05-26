//ff:func feature=gen-ir type=generator control=sequence
//ff:what BuildMiddlewarePlan -- manifest + prepared.State → MiddlewarePlan 변환

package ir

import (
	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// BuildMiddlewarePlan inspects manifest + prepared.State to determine
// which middlewares are active and populates their configuration. Pointer
// fields in the returned plan are nil when the corresponding middleware
// is inactive.
func BuildMiddlewarePlan(fs *yongol.Fullstack, ps *prepared.State) *MiddlewarePlan {
	plan := &MiddlewarePlan{
		RequestID:  true,
		Prometheus: prometheusEnabled(fs),
	}

	plan.BodyLimit = buildBodyLimitConfig(fs)
	plan.RateLimit = buildRateLimitConfig(fs)
	plan.CSRF = buildCSRFConfig(ps)
	plan.BearerAuth = buildBearerAuthConfig(ps)
	plan.ErrorEnvelope = buildErrorEnvelopeConfig(fs)
	plan.RequestValidator = buildRequestValidatorConfig()
	plan.SecurityHeaders = buildSecurityHeadersConfig(fs)

	return plan
}
