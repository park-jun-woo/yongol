//ff:func feature=gen-gogin type=generator control=iteration dimension=1 topic=rate-limit
//ff:what blockRateLimit — manifest.backend.rate_limit → RouteRateLimit 미들웨어 등록

package boot

import (
	"fmt"
	"sort"
	"time"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// blockRateLimit produces the route-aware rate limiting middleware
// registration. Active only when manifest.backend.rate_limit has entries
// and at least one operationId maps to an OpenAPI route.
//
// The block emits a rules map (route → RateLimitRule) and registers
// middleware.RouteRateLimit(rules) on the router. Placement: after
// blockBodyLimit and before blockRegisterHandlers so the rate limiter
// runs before the handler but after body size enforcement.
func blockRateLimit(fs *yongol.Fullstack, modulePath string) MainBlock {
	if fs == nil || fs.Manifest == nil || len(fs.Manifest.Backend.RateLimit) == 0 {
		return MainBlock{Name: "rate-limit"}
	}

	opToRoute := buildOperationRouteIndex(fs)
	type resolvedRule struct {
		route  string
		rate   int
		period time.Duration
		key    string
	}
	var rules []resolvedRule
	for opID, entry := range fs.Manifest.Backend.RateLimit {
		route, ok := opToRoute[opID]
		if !ok {
			continue
		}
		period, err := time.ParseDuration(entry.Period)
		if err != nil {
			continue
		}
		key := entry.Key
		if key == "" {
			key = "ip"
		}
		rules = append(rules, resolvedRule{
			route:  route,
			rate:   entry.Rate,
			period: period,
			key:    key,
		})
	}
	if len(rules) == 0 {
		return MainBlock{Name: "rate-limit"}
	}

	sort.Slice(rules, func(i, j int) bool { return rules[i].route < rules[j].route })

	var lines []string
	lines = append(lines, "rateLimitRules := map[string]middleware.RateLimitRule{")
	for _, r := range rules {
		lines = append(lines, fmt.Sprintf("\t%q: {Rate: %d, Period: %s, Key: %q},",
			r.route, r.rate, durationLiteral(r.period), r.key))
	}
	lines = append(lines, "}")
	lines = append(lines, "r.Use(middleware.RouteRateLimit(rateLimitRules))")

	return MainBlock{
		Name: "rate-limit",
		Imports: []string{
			fmt.Sprintf(`"%s/internal/middleware"`, modulePath),
			`"time"`,
		},
		Lines: lines,
	}
}
