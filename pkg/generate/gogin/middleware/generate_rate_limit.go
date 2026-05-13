//ff:func feature=gen-gogin type=generator control=sequence topic=rate-limit
//ff:what GenerateRateLimit — internal/middleware/fixed_rate_limit*.go 기록 (1 file 1 func)

package middleware

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// GenerateRateLimit emits internal/middleware/fixed_rate_limit.go and
// fixed_rate_limit_key.go — split so each file carries one func (filefunc F1).
func GenerateRateLimit(fs *yongol.Fullstack, artifactsDir string) error {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	if fs.Manifest.Backend.Module == "" {
		return nil
	}
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")

	files := renderRateLimitSources()
	if err := writeFiles(mwDir, files); err != nil {
		return fmt.Errorf("write rate_limit: %w", err)
	}
	return nil
}
