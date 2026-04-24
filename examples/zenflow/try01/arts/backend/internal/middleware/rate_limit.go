//ff:func feature=runtime-middleware type=util control=sequence topic=rate-limit
//ff:what FixedRateLimit — 특정 infra endpoint (/auth/refresh 등) 전용 하드코딩 rate 가드
//ff:checked llm=yongol-gen hash=6f9891af

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
