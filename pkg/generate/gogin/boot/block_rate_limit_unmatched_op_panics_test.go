//ff:func feature=gen-gogin type=test control=sequence topic=rate-limit
//ff:what blockRateLimit — 라우트 미매핑 operationId 는 SEC-05 가 선차단하므로 codegen 도달 시 방어 panic
package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// A rate_limit operationId with no OpenAPI route must never reach codegen —
// SEC-05 (validate) blocks it first. If it does reach here, blockRateLimit
// panics (defensive guard) rather than silently dropping the rate limiter
// (BUG-115).
func TestBlockRateLimit_UnmatchedOpPanics(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{RateLimit: pmanifest.RateLimitConfig{
			"Missing": {Rate: 5, Period: "1m"},
		}},
	}}
	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("expected panic on unmatched operationId, got none")
		}
		if msg, ok := r.(string); !ok || !strings.Contains(msg, "SEC-05") {
			t.Errorf("panic must reference SEC-05, got %v", r)
		}
	}()
	blockRateLimit(fs, "example.com/zenflow")
}
