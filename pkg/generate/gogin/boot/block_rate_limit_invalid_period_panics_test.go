//ff:func feature=gen-gogin type=test control=sequence topic=rate-limit
//ff:what blockRateLimit — 파싱 불가 period 는 C-10 이 선차단하므로 codegen 도달 시 방어 panic
package boot

import (
	"strings"
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// An unparseable period must never reach codegen — C-10 (validate) blocks
// it first. If it does reach here, blockRateLimit panics (defensive guard)
// rather than silently dropping the rate limiter (BUG-115).
func TestBlockRateLimit_InvalidPeriodPanics(t *testing.T) {
	doc := buildDoc([]opSpec{{path: "/login", method: "POST", opID: "Login"}}, false)
	fs := &yongol.Fullstack{
		OpenAPIDoc: doc,
		Manifest: &pmanifest.ProjectConfig{
			Backend: pmanifest.Backend{RateLimit: pmanifest.RateLimitConfig{
				"Login": {Rate: 5, Period: "not-a-duration"},
			}},
		},
	}
	defer func() {
		r := recover()
		if r == nil {
			t.Fatalf("expected panic on unparseable period, got none")
		}
		if msg, ok := r.(string); !ok || !strings.Contains(msg, "C-10") {
			t.Errorf("panic must reference C-10, got %v", r)
		}
	}()
	blockRateLimit(fs, "example.com/zenflow")
}
