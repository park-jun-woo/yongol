//ff:func feature=gen-gogin type=test control=sequence topic=rate-limit
//ff:what blockRateLimit — manifest.backend.rate_limit → RouteRateLimit 미들웨어 등록
package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockRateLimit_UnmatchedOpInert(t *testing.T) {
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{
		Backend: pmanifest.Backend{RateLimit: pmanifest.RateLimitConfig{
			"Missing": {Rate: 5, Period: "1m"},
		}},
	}}
	block := blockRateLimit(fs, "example.com/zenflow")
	if len(block.Lines) != 0 {
		t.Errorf("op with no matching route must yield inert block, got %v", block.Lines)
	}
}
