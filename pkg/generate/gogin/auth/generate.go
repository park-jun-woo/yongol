//ff:func feature=gen-gogin type=command control=sequence
//ff:what Generate — UserClaim struct + BearerAuth 미들웨어 생성 (JWT 런타임은 ssac/pkg/auth)

package auth

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/prepared"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Generate produces internal/model/user_claim.go and
// internal/middleware/bearerauth.go from manifest.backend.auth.claims.
// Skipped when claims is empty.
//
// Phase001 UserClaimUnification — the project-local `internal/auth` package
// is no longer generated. JWT Issue/Verify/Refresh and password hashing live
// in ssac/pkg/auth (shared runtime); the claim payload type is the project-
// local `model.UserClaim` generated from the manifest claim definitions.
// Any leftover `internal/auth/` directory from previous yongol versions is
// removed wholesale by cleanStaleAuthFiles.
func Generate(fs *yongol.Fullstack, p prepared.State, artifactsDir string) error {
	if !p.Auth.Present || p.Auth.Raw == nil || len(p.Auth.Raw.Claims) == 0 {
		return nil
	}
	fields := parseClaims(p.Auth.Raw.Claims)
	modulePath := fs.Manifest.Backend.Module

	if err := generateUserClaim(artifactsDir, fields); err != nil {
		return fmt.Errorf("user_claim: %w", err)
	}
	if fs.IsDomained() {
		// Phase008 — emit one strict middleware per domain (BearerAuthStrict
		// <Title>/CookieAuthStrict<Title>) by resolved auth_mode; the shared
		// auth_mode.go/extract_token.go are not written in domain mode (§3c/§3d).
		if err := generateDomainAuth(fs, p, artifactsDir, modulePath); err != nil {
			return fmt.Errorf("domain_auth: %w", err)
		}
	} else if err := generateBearerAuth(artifactsDir, modulePath, fields, p.Auth.Mode); err != nil {
		return fmt.Errorf("bearer_auth: %w", err)
	}
	// Phase001 UserClaimUnification — the whole internal/auth/ directory is
	// vestigial (previous versions emitted claim.go / reexport.go / issue_token.go
	// / refresh_token.go / verify_token.go / refresh_store.go / refresh_handler.go
	// there). Remove it in full on every regeneration so repeat `yongol
	// generate` runs against an older arts/ tree land in a clean state.
	if err := cleanStaleAuthFiles(artifactsDir); err != nil {
		return fmt.Errorf("clean stale auth files: %w", err)
	}
	return nil
}
