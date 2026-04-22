//ff:func feature=gen-gogin type=generator control=sequence
//ff:what GenerateBodyLimit — internal/middleware/body_limit.go 기록

package middleware

import (
	"fmt"
	"os"
	"path/filepath"
)

// GenerateBodyLimit emits internal/middleware/body_limit.go containing
// BodyLimit, MultipartLimit, OverrideBodyLimit gin middlewares. The source
// is static (no manifest-driven codegen) — runtime values are injected in
// main.go via BodyLimit(sizeVar) calls.
func GenerateBodyLimit(artifactsDir string) error {
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")
	if err := os.MkdirAll(mwDir, 0o755); err != nil {
		return fmt.Errorf("mkdir middleware: %w", err)
	}
	path := filepath.Join(mwDir, "body_limit.go")
	return os.WriteFile(path, []byte(bodyLimitSource), 0o644)
}
