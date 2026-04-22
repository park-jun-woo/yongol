//ff:type feature=gen-gogin type=generator
//ff:what rateLimitSourceTemplate — FixedRateLimit 헬퍼만 방출 (비즈니스 결합 가드 전용)

package middleware

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// rateLimitSourceTemplate carries the verbatim Go source for
// internal/middleware/rate_limit.go with __MODULE__ replaced by the
// manifest's backend.module. Scope after Phase006 deprecation:
//
//	FixedRateLimit(name, keyAxis, rate, period) gin.HandlerFunc
//
// Global rate limit + per-operationId strict rules + memory/redis/postgres
// store abstraction were all retired in favour of CDN/WAF/Gateway layers
// (see plans/deprecated/Phase006-DeprecateAppLayerRateLimit.md). Only the
// tight business-logic guards that the generated server mounts outside
// OpenAPI (e.g. /auth/refresh 10 rpm/IP) live here now.
const rateLimitSourceTemplate = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=rate-limit
//` + `ff:what FixedRateLimit — 특정 infra endpoint (/auth/refresh 등) 전용 하드코딩 rate 가드

package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// FixedRateLimit is a convenience wrapper for infrastructure endpoints that
// are mounted outside the oapi-codegen tree (e.g. /auth/refresh). Builds a
// fresh memory-backed ulule limiter on demand from the provided rate +
// window and rejects excess requests with a 429 + Retry-After header.
//
// keyAxis is currently "ip" only — other axes (email / user_id) need
// access to request body or authenticated context and are therefore
// handled by dedicated guards rather than this helper.
func FixedRateLimit(name string, keyAxis string, rate int, period time.Duration) gin.HandlerFunc {
	l := limiter.New(memory.NewStore(), limiter.Rate{Period: period, Limit: int64(rate)})
	return func(c *gin.Context) {
		key := fixedRateLimitKey(c, keyAxis)
		if key == "" {
			c.Next()
			return
		}
		lctx, err := l.Get(c, fmt.Sprintf("fixed:%s:%s:%s", name, keyAxis, key))
		if err != nil {
			// memory store never errors; fall through on unexpected failure.
			c.Next()
			return
		}
		if lctx.Reached {
			retry := lctx.Reset - time.Now().Unix()
			if retry < 0 {
				retry = 0
			}
			c.Header("X-RateLimit-Remaining", "0")
			c.Header("Retry-After", fmt.Sprint(retry))
			WriteEnvelopeWithContext(c,
				http.StatusTooManyRequests,
				"rate_limit_exceeded",
				"",
				map[string]interface{}{"retry_after": lctx.Reset},
			)
			c.Abort()
			return
		}
		c.Header("X-RateLimit-Remaining", fmt.Sprint(lctx.Remaining))
		c.Next()
	}
}

// fixedRateLimitKey extracts the rate-limit key axis from the gin context.
// Only "ip" is supported; other axes return "" so the guard is skipped —
// callers should not wire FixedRateLimit with an unsupported axis.
func fixedRateLimitKey(c *gin.Context, key string) string {
	switch key {
	case "ip":
		return c.ClientIP()
	}
	return ""
}
`

// renderRateLimitSource substitutes __MODULE__ with the actual module path.
func renderRateLimitSource(modulePath string) string {
	return strings.ReplaceAll(rateLimitSourceTemplate, "__MODULE__", modulePath)
}

// GenerateRateLimit emits internal/middleware/rate_limit.go containing only
// FixedRateLimit — used by Phase002 /auth/refresh guard (block_auth_refresh).
// The pre-deprecation rate_limit_store.go is no longer produced; callers
// that need gateway-layer rate limiting should configure their CDN/WAF or
// API gateway (see plans/deprecated/Phase006-DeprecateAppLayerRateLimit.md).
func GenerateRateLimit(fs *yongol.Fullstack, artifactsDir string) error {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	modulePath := fs.Manifest.Backend.Module
	if modulePath == "" {
		return nil
	}
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")
	if err := os.MkdirAll(mwDir, 0o755); err != nil {
		return fmt.Errorf("mkdir middleware: %w", err)
	}
	rlPath := filepath.Join(mwDir, "rate_limit.go")
	if err := os.WriteFile(rlPath, []byte(renderRateLimitSource(modulePath)), 0o644); err != nil {
		return fmt.Errorf("write rate_limit.go: %w", err)
	}
	return nil
}
