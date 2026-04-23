//ff:func feature=gen-gogin type=generator control=sequence
//ff:what GenerateRateLimit — internal/middleware/rate_limit.go 기록 (FixedRateLimit 만 방출)

package middleware

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// GenerateRateLimit emits internal/middleware/rate_limit.go containing only
// FixedRateLimit — used by Phase002 /auth/refresh guard (block_auth_refresh).
// The pre-deprecation rate_limit_store.go is no longer produced; callers
// that need gateway-layer rate limiting should configure their CDN/WAF or
// API gateway (see plans/deprecated/Phase006-DeprecateAppLayerRateLimit.md).
func GenerateRateLimit(fs *yongol.Fullstack, artifactsDir string) error {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	modulePath := fs.Manifest.Backend.Module
	if modulePath == "" {
		return nil
	}
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")
	if err := os.MkdirAll(mwDir, 0o755); err != nil {
		return fmt.Errorf("mkdir middleware: %w", err)
	}
	rlPath := filepath.Join(mwDir, "rate_limit.go")
	if err := os.WriteFile(rlPath, []byte(renderRateLimitSource(modulePath)), 0o644); err != nil {
		return fmt.Errorf("write rate_limit.go: %w", err)
	}
	return nil
}
