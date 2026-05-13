//ff:func feature=gen-gogin type=generator control=sequence topic=error-envelope
//ff:what GenerateErrorEnvelope — internal/middleware/error_envelope*.go 기록 (1 file 1 func)

package middleware

import (
	"fmt"
	"path/filepath"
)

// GenerateErrorEnvelope emits the error envelope middleware files split so
// each file carries one func or type (filefunc F1/F2): error_envelope.go
// (type), default_code_for.go, default_message_for.go, write_envelope.go,
// write_envelope_with_context.go, error_envelope_middleware.go.
func GenerateErrorEnvelope(artifactsDir string) error {
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")
	files := map[string]string{
		"error_envelope.go":              errorEnvelopeTypeSource,
		"default_code_for.go":            defaultCodeForSource,
		"default_message_for.go":         defaultMessageForSource,
		"write_envelope.go":              writeEnvelopeSource,
		"write_envelope_with_context.go": writeEnvelopeWithContextSource,
		"error_envelope_middleware.go":   errorEnvelopeMiddlewareSource,
	}
	if err := writeFiles(mwDir, files); err != nil {
		return fmt.Errorf("write error_envelope: %w", err)
	}
	return nil
}
