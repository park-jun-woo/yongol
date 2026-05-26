//ff:type feature=gen-ir type=model
//ff:what RateLimitRule -- 단일 라우트의 rate limit 항목

package ir

import "time"

// RateLimitRule is a single per-route rate limit entry.
type RateLimitRule struct {
	// Route is the HTTP route path (e.g. "/courses/:id").
	Route string

	// Rate is the maximum number of requests allowed per Period.
	Rate int

	// Period is the time window for the rate limit.
	Period time.Duration

	// Key determines how requests are bucketed ("ip" or "user").
	Key string
}
