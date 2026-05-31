//ff:func feature=gen-gogin type=test control=iteration dimension=1
//ff:what TestMiddlewareExtras — csrf/prometheus/rate-limit/request-id/writeFiles 유틸 검증
package middleware

import (
	"testing"
)

func TestRenderRateLimitSources(t *testing.T) {
	m := renderRateLimitSources()
	for _, k := range []string{"fixed_rate_limit.go", "fixed_rate_limit_key.go", "route_rate_limit.go", "route_rate_limit_key.go"} {
		if _, ok := m[k]; !ok {
			t.Errorf("missing rate-limit source %q", k)
		}
	}
}
