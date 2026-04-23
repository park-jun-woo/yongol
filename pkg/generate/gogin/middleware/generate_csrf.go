//ff:func feature=gen-gogin type=generator control=sequence topic=csrf
//ff:what GenerateCsrf — internal/middleware/csrf.go 기록 (auth.mode=cookie|hybrid 시)

package middleware

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
)

// GenerateCsrf emits internal/middleware/csrf.go when the resolved auth
// mode is "cookie" or "hybrid" AND csrf.enabled=true. On bearer-only
// projects (default) the file is skipped entirely — no dead code in the
// artifacts tree.
//
// Phase005 remains dormant: with the default bearer mode active, this
// function early-returns without writing the file. Once Phase020 lands
// the cookie-session infrastructure, flipping auth.mode will auto-emit
// the middleware source.
func GenerateCsrf(a prepared.Auth, artifactsDir string) error {
	if !csrfActive(a) {
		return nil
	}
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")
	if err := os.MkdirAll(mwDir, 0o755); err != nil {
		return fmt.Errorf("mkdir middleware: %w", err)
	}
	path := filepath.Join(mwDir, "csrf.go")
	if err := os.WriteFile(path, []byte(csrfSourceTemplate), 0o644); err != nil {
		return fmt.Errorf("write csrf.go: %w", err)
	}
	return nil
}
