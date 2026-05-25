//ff:type feature=gen-gogin type=generator topic=rate-limit
//ff:what routeRateLimitSource — RouteRateLimit 미들웨어 소스 (경로 기반 rate limit 분기)

package middleware

// routeRateLimitSource is the source for internal/middleware/route_rate_limit.go.
const routeRateLimitSource = `//` + `ff:func feature=runtime-middleware type=util control=sequence topic=rate-limit
//` + `ff:what RouteRateLimit — 라우트별 rate limit 분기 미들웨어 (manifest rate_limit 기반)

package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ulule/limiter/v3"
	"github.com/ulule/limiter/v3/drivers/store/memory"
)

// RateLimitRule holds the configuration for a single route rate limit.
type RateLimitRule struct {
	Rate   int
	Period time.Duration
	Key    string
}

// RouteRateLimit applies per-route rate limiting based on the gin route
// template (c.FullPath()). Routes not present in the map are unaffected.
// Each rule gets its own memory-backed limiter instance.
func RouteRateLimit(rules map[string]RateLimitRule) gin.HandlerFunc {
	limiters := make(map[string]*limiter.Limiter, len(rules))
	for route, rule := range rules {
		limiters[route] = limiter.New(
			memory.NewStore(),
			limiter.Rate{Period: rule.Period, Limit: int64(rule.Rate)},
		)
	}
	return func(c *gin.Context) {
		route := c.Request.Method + " " + c.FullPath()
		l, ok := limiters[route]
		if !ok {
			c.Next()
			return
		}
		rule := rules[route]
		key := routeRateLimitKey(c, rule.Key)
		if key == "" {
			c.Next()
			return
		}
		lctx, err := l.Get(c, fmt.Sprintf("route:%s:%s", route, key))
		if err != nil {
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
`

// routeRateLimitKeySource is the source for internal/middleware/route_rate_limit_key.go.
const routeRateLimitKeySource = `//` + `ff:func feature=runtime-middleware type=util control=selection topic=rate-limit
//` + `ff:what routeRateLimitKey — route rate-limit key axis(ip) 추출

package middleware

import "github.com/gin-gonic/gin"

// routeRateLimitKey extracts the rate-limit key from the gin context.
// Only "ip" is supported; other axes return "" so the guard is skipped.
func routeRateLimitKey(c *gin.Context, key string) string {
	switch key {
	case "ip":
		return c.ClientIP()
	}
	return ""
}
`
