//ff:func feature=gen-gogin type=command control=sequence
//ff:what Generate — Claim struct + reexport + BearerAuth 미들웨어 생성 (JWT 런타임은 ssac/pkg/auth)

package auth

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Generate produces internal/auth/ (Claim struct + ssac/pkg/auth reexport),
// internal/middleware/bearerauth.go, and internal/model/current_user.go.
// All derived from manifest.backend.auth.claims. Skipped when claims is empty.
//
// JWT 3함수 (IssueToken / RefreshToken / VerifyToken) and refresh-token
// rotation (RefreshStore / RefreshHandler / RefreshTokensDDL) live in
// ssac/pkg/auth as of Phase003. yongol only generates the project-local
// `Claim` struct whose JSON tags align with manifest claim keys.
func Generate(fs *yongol.Fullstack, artifactsDir string) error {
	if fs.Manifest == nil || fs.Manifest.Backend.Auth == nil || len(fs.Manifest.Backend.Auth.Claims) == 0 {
		return nil
	}
	fields := parseClaims(fs.Manifest.Backend.Auth.Claims)
	modulePath := fs.Manifest.Backend.Module

	authDir := filepath.Join(artifactsDir, "backend", "internal", "auth")
	if err := os.MkdirAll(authDir, 0o755); err != nil {
		return fmt.Errorf("mkdir auth: %w", err)
	}

	if err := generateCurrentUser(artifactsDir, fields); err != nil {
		return fmt.Errorf("current_user: %w", err)
	}
	if err := generateClaim(authDir, fields); err != nil {
		return fmt.Errorf("claim: %w", err)
	}
	if err := generateReexport(authDir); err != nil {
		return fmt.Errorf("reexport: %w", err)
	}
	defaultMode := fs.Manifest.Backend.Auth.ResolvedMode()
	if err := generateBearerAuth(artifactsDir, modulePath, fields, defaultMode); err != nil {
		return fmt.Errorf("bearer_auth: %w", err)
	}
	// Phase003 — previous yongol versions emitted issue_token.go /
	// refresh_token.go / verify_token.go / refresh_store.go / refresh_handler.go
	// into internal/auth. Those files now live in ssac/pkg/auth, so sweep
	// them away after each regeneration to avoid duplicate-symbol build
	// errors on repeat runs of `yongol generate`.
	if err := cleanStaleAuthFiles(authDir); err != nil {
		return fmt.Errorf("clean stale auth files: %w", err)
	}
	return nil
}
