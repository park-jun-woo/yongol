//ff:func feature=gen-gogin type=test control=sequence
//ff:what TestMiddlewareExtras — csrf/prometheus/rate-limit/request-id/writeFiles 유틸 검증
package middleware

import (
	"testing"

	pmanifest "github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGenerateRateLimitWriteError(t *testing.T) {
	arts := makeMiddlewareDirCollision(t)
	fs := &yongol.Fullstack{Manifest: &pmanifest.ProjectConfig{Backend: pmanifest.Backend{Module: "example.com/app"}}}
	if err := GenerateRateLimit(fs, arts); err == nil {
		t.Errorf("expected write rate_limit error, got nil")
	}
}
