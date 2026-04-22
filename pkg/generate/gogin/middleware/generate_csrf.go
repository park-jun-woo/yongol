//ff:func feature=gen-gogin type=generator control=sequence topic=csrf
//ff:what GenerateCsrf — internal/middleware/csrf.go 기록 (auth.mode=cookie|hybrid 시)

package middleware

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// GenerateCsrf emits internal/middleware/csrf.go when manifest declares
// backend.auth.mode="cookie" or "hybrid" AND csrf.enabled=true. On
// bearer-only projects (default) the file is skipped entirely — no dead
// code in the artifacts tree.
//
// Phase005 remains dormant: with the default bearer mode active, this
// function early-returns without writing the file. Once Phase020 lands
// the cookie-session infrastructure, flipping auth.mode will auto-emit
// the middleware source.
func GenerateCsrf(fs *yongol.Fullstack, artifactsDir string) error {
	if !csrfActive(fs) {
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

// csrfActive mirrors the boot.blockCsrf Active condition so the
// middleware file is emitted on the same trigger. Kept local to avoid
// import cycles (middleware → boot).
func csrfActive(fs *yongol.Fullstack) bool {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return false
	}
	a := fs.Manifest.Backend.Auth
	mode := a.Mode
	if mode != "cookie" && mode != "hybrid" {
		return false
	}
	if a.Csrf == nil {
		// Default on for cookie/hybrid modes — SEC-201 rejects the
		// explicit-false combination at validate time; reaching codegen
		// with nil Csrf means "accept defaults, enabled".
		return true
	}
	return a.Csrf.Enabled
}
