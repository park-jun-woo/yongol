//ff:func feature=gen-gogin type=test control=iteration dimension=1 topic=rate-limit
//ff:what blockRateLimit — manifest.backend.rate_limit → RouteRateLimit 미들웨어 등록

package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBlockRateLimit_Empty(t *testing.T) {
	block := blockRateLimit(&yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{}}, "example.com/zenflow")
	if len(block.Lines) != 0 {
		t.Errorf("no rate_limit must produce inert block, got %v", block.Lines)
	}
}

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

func TestBlockRateLimit_NilFullstack(t *testing.T) {
	block := blockRateLimit(nil, "example.com/zenflow")
	if len(block.Lines) != 0 {
		t.Errorf("nil Fullstack must yield inert block, got %v", block.Lines)
	}
}

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

func TestBlockRateLimit_MultiRouteSorted(t *testing.T) {
	doc := buildDoc([]opSpec{
		{path: "/login", method: "POST", opID: "Login"},
		{path: "/signup", method: "POST", opID: "Signup"},
	}, false)
	fs := &yongol.Fullstack{
		OpenAPIDoc: doc,
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{RateLimit: pmanifest.RateLimitConfig{
				"Login":  {Rate: 5, Period: "1m"},
				"Signup": {Rate: 3, Period: "1m"},
			}},
		},
	}
	block := blockRateLimit(fs, "example.com/zenflow")
	body := strings.Join(block.Lines, "\n")
	loginIdx := strings.Index(body, "POST /login")
	signupIdx := strings.Index(body, "POST /signup")
	if loginIdx < 0 || signupIdx < 0 {
		t.Fatalf("both routes must be present, got:\n%s", body)
	}
	if loginIdx > signupIdx {
		t.Errorf("routes must be sorted ascending, got:\n%s", body)
	}
}

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
