//ff:func feature=gen-gogin type=generator control=sequence
//ff:what generateRequestValidator — OpenAPI 스펙 기반 런타임 validation 미들웨어 생성

package middleware

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Generate writes internal/middleware/request_validator.go and copies
// specs/api/openapi.yaml next to it so `go:embed` can bundle the spec.
// The generated middleware runs kin-openapi ValidateRequest on every request
// (except /health /ready /metrics prefixes) and rejects payloads that violate
// OpenAPI schema constraints (minLength, maximum, pattern, required, enum…).
func Generate(fs *yongol.Fullstack, p prepared.State, artifactsDir string) error {
	if fs == nil {
		return nil
	}
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")
	if err := os.MkdirAll(mwDir, 0o755); err != nil {
		return fmt.Errorf("mkdir middleware: %w", err)
	}

	// Remove stale combined files from previous runs (e.g. prometheus.go
	// replaced by prometheus_middleware.go + prometheus_handler.go).
	if err := removeStaleCombined(mwDir); err != nil {
		return fmt.Errorf("remove stale combined: %w", err)
	}

	// Copy specs/api/openapi.yaml → middleware/openapi.yaml for go:embed.
	src := filepath.Join(fs.SpecsDir, "api", "openapi.yaml")
	dst := filepath.Join(mwDir, "openapi.yaml")
	if err := copyFile(src, dst); err != nil {
		return fmt.Errorf("copy openapi.yaml: %w", err)
	}

	// Write request_validator.go
	if err := writeValidator(mwDir); err != nil {
		return fmt.Errorf("write validator: %w", err)
	}
	// Write body_limit.go (Phase003 DoS guard).
	if err := GenerateBodyLimit(artifactsDir); err != nil {
		return fmt.Errorf("write body_limit: %w", err)
	}
	// Write rate_limit.go + rate_limit_store.go (Phase006).
	if err := GenerateRateLimit(fs, artifactsDir); err != nil {
		return fmt.Errorf("write rate_limit: %w", err)
	}
	// Write prometheus.go (Phase008 observability).
	if err := GeneratePrometheus(fs, artifactsDir); err != nil {
		return fmt.Errorf("write prometheus: %w", err)
	}
	// Write csrf.go (Phase005 — dormant unless auth.mode=cookie|hybrid).
	if err := GenerateCsrf(p.Auth, artifactsDir); err != nil {
		return fmt.Errorf("write csrf: %w", err)
	}
	// Write security_headers.go (Phase007 browser security headers).
	if err := GenerateSecurityHeaders(fs, artifactsDir); err != nil {
		return fmt.Errorf("write security_headers: %w", err)
	}
	// Write request_id.go (Phase004 correlation id).
	if err := GenerateRequestID(artifactsDir); err != nil {
		return fmt.Errorf("write request_id: %w", err)
	}
	// Write error_envelope.go (Phase004 canonical error body).
	if err := GenerateErrorEnvelope(artifactsDir); err != nil {
		return fmt.Errorf("write error_envelope: %w", err)
	}
	return nil
}
