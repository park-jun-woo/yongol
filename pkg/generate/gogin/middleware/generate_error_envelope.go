//ff:func feature=gen-gogin type=generator control=sequence topic=error-envelope
//ff:what GenerateErrorEnvelope — internal/middleware/error_envelope.go 기록 (Phase004)

package middleware

import (
	"fmt"
	"os"
	"path/filepath"
)

// GenerateErrorEnvelope emits internal/middleware/error_envelope.go
// containing the canonical ErrorEnvelope struct, status → code/message
// tables, and the ErrorEnvelopeMiddleware gin middleware. Runtime toggles
// (ExposeInternalError) are set in main.go via blockErrorEnvelope.
func GenerateErrorEnvelope(artifactsDir string) error {
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")
	if err := os.MkdirAll(mwDir, 0o755); err != nil {
		return fmt.Errorf("mkdir middleware: %w", err)
	}
	path := filepath.Join(mwDir, "error_envelope.go")
	return os.WriteFile(path, []byte(errorEnvelopeSource), 0o644)
}
