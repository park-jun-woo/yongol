//ff:type feature=gen-ir type=model
//ff:what MiddlewarePlan -- 미들웨어 활성화의 프레임워크 비의존 중간 표현

package ir

// MiddlewarePlan is the framework-agnostic intermediate representation of
// the middleware layer. Each field represents one middleware subsystem;
// pointer fields are nil when the subsystem is inactive, bool fields
// indicate unconditionally active subsystems.
//
// Backend renderers (gogin, nestjs, fastapi) consume this plan to produce
// framework-specific middleware registration code. The activation logic
// is pre-evaluated by BuildMiddlewarePlan.
type MiddlewarePlan struct {
	// RequestID is always active (ULID cost is negligible).
	RequestID bool

	// BodyLimit configures body/multipart size limits. nil when inactive.
	BodyLimit *BodyLimitConfig

	// RateLimit configures per-route rate limiting. nil when no rules exist.
	RateLimit *RateLimitConfig

	// CSRF configures double-submit cookie defense. nil only when auth is
	// absent or csrf is explicitly disabled — bearer builds carry it too
	// (BUG-116: runtime-gated, no-op until BACKEND_AUTH_MODE=cookie/hybrid).
	CSRF *CSRFConfig

	// Prometheus is true when metrics middleware is active (default true).
	Prometheus bool

	// BearerAuth configures JWT/cookie/hybrid auth middleware. nil when
	// no auth is declared.
	BearerAuth *BearerAuthConfig

	// ErrorEnvelope configures error response standardization. nil when
	// inactive (practically always present).
	ErrorEnvelope *ErrorEnvelopeConfig

	// RequestValidator configures OpenAPI constraint validation middleware.
	// nil when inactive (practically always present).
	RequestValidator *RequestValidatorConfig

	// SecurityHeaders configures browser security header middleware.
	// nil when explicitly disabled.
	SecurityHeaders *SecurityHeadersConfig
}
