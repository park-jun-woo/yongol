//ff:func feature=gen-gogin type=generator control=sequence topic=security-headers
//ff:what GenerateSecurityHeaders — internal/middleware/security_headers.go 기록

package middleware

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// GenerateSecurityHeaders emits internal/middleware/security_headers.go.
// The file is static — all runtime behavior is driven by a
// SecurityHeadersConfig value assembled in generated main.go (blockSecurityHeaders)
// from manifest.backend.security_headers + env overrides. When the manifest
// disables security headers (enabled=false) the file is still written so
// importing middleware.SecurityHeadersMiddleware never fails to compile;
// blockSecurityHeaders decides whether to call it.
func GenerateSecurityHeaders(fs *yongol.Fullstack, artifactsDir string) error {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	if fs.Manifest.Backend.Module == "" {
		return nil
	}
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")
	if err := os.MkdirAll(mwDir, 0o755); err != nil {
		return fmt.Errorf("mkdir middleware: %w", err)
	}
	path := filepath.Join(mwDir, "security_headers.go")
	if err := os.WriteFile(path, []byte(securityHeadersSource), 0o644); err != nil {
		return fmt.Errorf("write security_headers.go: %w", err)
	}
	return nil
}
