//ff:func feature=gen-gogin type=generator control=sequence topic=request-id
//ff:what GenerateRequestID — internal/middleware/request_id*.go 기록 (1 file 1 func)

package middleware

import (
	"fmt"
	"path/filepath"
)

// GenerateRequestID emits request-id middleware files split so each carries
// one func or type (filefunc F1/F2): request_id_type.go (type/constants),
// request_id.go (func), request_id_from_context.go,
// request_id_from_std_context.go, generate_request_id_ulid.go,
// sanitize_upstream_id.go.
func GenerateRequestID(artifactsDir string) error {
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")
	files := map[string]string{
		"request_id_type.go":             requestIDTypeSource,
		"request_id.go":                  requestIDMainSource,
		"request_id_from_context.go":     requestIDFromContextSource,
		"request_id_from_std_context.go": requestIDFromStdContextSource,
		"generate_request_id_ulid.go":    generateRequestIDSource,
		"sanitize_upstream_id.go":        sanitizeUpstreamIDSource,
	}
	if err := writeFiles(mwDir, files); err != nil {
		return fmt.Errorf("write request_id: %w", err)
	}
	return nil
}
