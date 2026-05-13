//ff:func feature=gen-gogin type=generator control=sequence topic=dos-guard
//ff:what GenerateBodyLimit — internal/middleware/body_limit*.go 기록 (1 file 1 func)

package middleware

import (
	"fmt"
	"path/filepath"
)

// GenerateBodyLimit emits body-limit middleware files split so each carries
// one func (filefunc F1): body_limit.go, multipart_limit.go,
// override_body_limit.go, apply_override.go, respond_if_body_too_large.go.
func GenerateBodyLimit(artifactsDir string) error {
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")
	files := map[string]string{
		"body_limit.go":                bodyLimitSource,
		"multipart_limit.go":           multipartLimitSource,
		"override_body_limit.go":       overrideBodyLimitSource,
		"apply_override.go":            applyOverrideSource,
		"respond_if_body_too_large.go": respondIfBodyTooLargeSource,
	}
	if err := writeFiles(mwDir, files); err != nil {
		return fmt.Errorf("write body_limit: %w", err)
	}
	return nil
}
