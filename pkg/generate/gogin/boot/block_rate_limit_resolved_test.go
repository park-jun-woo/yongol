//ff:func feature=gen-gogin type=test control=sequence topic=rate-limit
//ff:what blockRateLimit — manifest.backend.rate_limit → RouteRateLimit 미들웨어 등록
package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockRateLimit_Resolved(t *testing.T) {
	doc := buildDoc([]opSpec{{path: "/login", method: "POST", opID: "Login"}}, false)
	fs := &yongol.Fullstack{
		OpenAPIDoc: doc,
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{RateLimit: pmanifest.RateLimitConfig{
				"Login": {Rate: 5, Period: "1m"}, // no key → defaults to "ip"
			}},
		},
	}
	block := blockRateLimit(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	if !strings.Contains(body, "rateLimitRules := map[string]middleware.RateLimitRule{") {
		t.Errorf("must emit rules map, got:\n%s", body)
	}
	if !strings.Contains(body, `"POST /login": {Rate: 5, Period: time.Duration(60000000000), Key: "ip"},`) {
		t.Errorf("rule entry wrong, got:\n%s", body)
	}
	if !strings.Contains(body, "r.Use(middleware.RouteRateLimit(rateLimitRules))") {
		t.Errorf("must register RouteRateLimit, got:\n%s", body)
	}
}
