//ff:func feature=gen-gogin type=test control=sequence topic=rate-limit
//ff:what blockRateLimit — manifest.backend.rate_limit → RouteRateLimit 미들웨어 등록
package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
