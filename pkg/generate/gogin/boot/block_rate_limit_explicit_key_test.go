//ff:func feature=gen-gogin type=test control=sequence topic=rate-limit
//ff:what blockRateLimit — manifest.backend.rate_limit → RouteRateLimit 미들웨어 등록
package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockRateLimit_ExplicitKey(t *testing.T) {
	doc := buildDoc([]opSpec{{path: "/login", method: "POST", opID: "Login"}}, false)
	fs := &yongol.Fullstack{
		OpenAPIDoc: doc,
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{RateLimit: pmanifest.RateLimitConfig{
				"Login": {Rate: 5, Period: "1m", Key: "user"},
			}},
		},
	}
	block := blockRateLimit(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, `Key: "user"`) {
		t.Errorf("explicit key must be preserved, got:\n%s", body)
	}
}
