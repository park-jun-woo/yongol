//ff:func feature=gen-gogin type=test control=sequence topic=rate-limit
//ff:what blockRateLimit — manifest.backend.rate_limit → RouteRateLimit 미들웨어 등록
package boot

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockRateLimit_InvalidPeriodInert(t *testing.T) {
	doc := buildDoc([]opSpec{{path: "/login", method: "POST", opID: "Login"}}, false)
	fs := &yongol.Fullstack{
		OpenAPIDoc: doc,
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{RateLimit: pmanifest.RateLimitConfig{
				"Login": {Rate: 5, Period: "not-a-duration"},
			}},
		},
	}
	block := blockRateLimit(fs, "example.com/zenflow")
	if len(block.Lines) != 0 {
		t.Errorf("invalid period must skip rule -> inert block, got %v", block.Lines)
	}
}
