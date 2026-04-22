//ff:func feature=gen-gogin type=generator control=sequence topic=request-id
//ff:what GenerateRequestID — internal/middleware/request_id.go 기록 (Phase004)

package middleware

import (
	"fmt"
	"os"
	"path/filepath"
)

// GenerateRequestID emits internal/middleware/request_id.go containing the
// RequestID gin middleware, ULID-based generator, and context helpers. The
// source is static — runtime policy (trust_upstream, header name) is passed
// in from main.go via blockRequestID.
func GenerateRequestID(artifactsDir string) error {
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")
	if err := os.MkdirAll(mwDir, 0o755); err != nil {
		return fmt.Errorf("mkdir middleware: %w", err)
	}
	path := filepath.Join(mwDir, "request_id.go")
	return os.WriteFile(path, []byte(requestIDSource), 0o644)
}
