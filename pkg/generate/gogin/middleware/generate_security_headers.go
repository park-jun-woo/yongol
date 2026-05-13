//ff:func feature=gen-gogin type=generator control=sequence topic=security-headers
//ff:what GenerateSecurityHeaders — internal/middleware/security_headers*.go 기록 (1 file 1 func)

package middleware

import (
	"fmt"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// GenerateSecurityHeaders emits security headers middleware files split so
// each file carries one func or type (filefunc F1/F2).
func GenerateSecurityHeaders(fs *yongol.Fullstack, artifactsDir string) error {
	if fs == nil || fs.Manifest == nil {
		return nil
	}
	if fs.Manifest.Backend.Module == "" {
		return nil
	}
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")
	files := map[string]string{
		"security_headers_config.go":       securityHeadersConfigSource,
		"security_headers_middleware.go":   securityHeadersMiddlewareSource,
		"build_static_security_headers.go": buildStaticSecurityHeadersSource,
		"build_csp_header.go":              buildCSPHeaderSource,
		"build_csp_value.go":               buildCSPValueSource,
		"build_permissions_policy.go":      buildPermissionsPolicySource,
	}
	if err := writeFiles(mwDir, files); err != nil {
		return fmt.Errorf("write security_headers: %w", err)
	}
	return nil
}
