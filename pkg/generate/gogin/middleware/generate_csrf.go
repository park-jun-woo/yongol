//ff:func feature=gen-gogin type=generator control=sequence topic=csrf
//ff:what GenerateCsrf — internal/middleware/csrf.go 기록 (auth 선언 시, bearer 는 런타임 게이트로 no-op)

package middleware

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
)

// GenerateCsrf emits internal/middleware/csrf.go whenever auth is declared
// and CSRF is not explicitly disabled (csrfActive). On bearer builds the
// file is still emitted, but the runtime gate (csrfRuntimeActive) keeps it
// inert until BACKEND_AUTH_MODE selects cookie/hybrid.
//
// BUG-116 / Phase-B1 — the build-time resolved mode (a.Mode) is injected as
// the csrfAuthMode() fallback so that, with BACKEND_AUTH_MODE unset, the
// middleware matches the manifest default (cookie/hybrid → active, bearer →
// no-op). Only when an operator overrides the env does the gate flip.
func GenerateCsrf(a prepared.Auth, artifactsDir string) error {
	if !csrfActive(a) {
		return nil
	}
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")
	if err := os.MkdirAll(mwDir, 0o755); err != nil {
		return fmt.Errorf("mkdir middleware: %w", err)
	}
	path := filepath.Join(mwDir, "csrf.go")
	source := fmt.Sprintf(csrfSourceTemplate, a.Mode)
	if err := os.WriteFile(path, []byte(source), 0o644); err != nil {
		return fmt.Errorf("write csrf.go: %w", err)
	}
	return nil
}
