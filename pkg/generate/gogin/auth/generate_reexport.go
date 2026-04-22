//ff:func feature=gen-gogin type=generator control=sequence
//ff:what generateReexport — writes internal/auth/reexport.go (full re-export of ssac/pkg/auth)

package auth

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
)

// generateReexport writes internal/auth/reexport.go. After Phase003 the
// entire auth runtime lives in ssac/pkg/auth, so this file is a single
// alias surface that lets project code import "internal/auth" uniformly.
//
// Re-exported symbols:
//   - Password hashing — HashPassword / VerifyPassword / DummyHash / GenerateResetToken
//   - JWT issue/verify — Configure / IssueToken / RefreshToken / VerifyToken
//   - Refresh rotation — RefreshStore / RefreshHandler / RefreshTokensDDL /
//     ClaimMatcher / ErrRefreshTokenNotFound / ErrRefreshTokenReused
//   - Auth endpoints (Phase009) — RefreshRotate / Logout (standard SSaC path)
//
// Type aliases use `=` so `auth.IssueTokenRequest` in callers is the exact
// same type as `ssac/pkg/auth.IssueTokenRequest`.
func generateReexport(authDir string) error {
	header := ffannot.EmitAnnotationBlock(ffannot.Block{
		Type: ffannot.TypeAnnot{Feature: "auth", Type: "accessor"},
		What: "reexport — full re-export of ssac/pkg/auth (password + JWT + refresh rotation)",
	})
	src := header + `package auth

import pkgauth "github.com/park-jun-woo/ssac/pkg/auth"

// Password hashing (Phase010 carry-over).
var HashPassword = pkgauth.HashPassword
var VerifyPassword = pkgauth.VerifyPassword
var GenerateResetToken = pkgauth.GenerateResetToken

// DummyHash — bcrypt timing-safe constant re-export.
const DummyHash = pkgauth.DummyHash

type HashPasswordRequest = pkgauth.HashPasswordRequest
type HashPasswordResponse = pkgauth.HashPasswordResponse
type VerifyPasswordRequest = pkgauth.VerifyPasswordRequest
type VerifyPasswordResponse = pkgauth.VerifyPasswordResponse
type GenerateResetTokenRequest = pkgauth.GenerateResetTokenRequest
type GenerateResetTokenResponse = pkgauth.GenerateResetTokenResponse

// JWT issue/verify/refresh (Phase003 — moved to ssac/pkg/auth).
var Configure = pkgauth.Configure
var IssueToken = pkgauth.IssueToken
var RefreshToken = pkgauth.RefreshToken
var VerifyToken = pkgauth.VerifyToken

type Config = pkgauth.Config
type IssueTokenRequest = pkgauth.IssueTokenRequest
type IssueTokenResponse = pkgauth.IssueTokenResponse
type RefreshTokenRequest = pkgauth.RefreshTokenRequest
type RefreshTokenResponse = pkgauth.RefreshTokenResponse
type VerifyTokenRequest = pkgauth.VerifyTokenRequest
type VerifyTokenResponse = pkgauth.VerifyTokenResponse

// Refresh rotation (Phase003 — moved to ssac/pkg/auth).
var RefreshHandler = pkgauth.RefreshHandler
var RefreshTokensDDL = pkgauth.RefreshTokensDDL
var ErrRefreshTokenNotFound = pkgauth.ErrRefreshTokenNotFound
var ErrRefreshTokenReused = pkgauth.ErrRefreshTokenReused

type RefreshStore = pkgauth.RefreshStore
type ClaimMatcher = pkgauth.ClaimMatcher

// Auth endpoints (Phase009 — standard SSaC path).
var RefreshRotate = pkgauth.RefreshRotate
var Logout = pkgauth.Logout

type RefreshRotateRequest = pkgauth.RefreshRotateRequest
type RefreshRotateResponse = pkgauth.RefreshRotateResponse
type LogoutRequest = pkgauth.LogoutRequest
type LogoutResponse = pkgauth.LogoutResponse

// Cookie session helpers (Phase020 — 2026 standard: HttpOnly + __Host- +
// SameSite). SetAuthCookies / ClearAuthCookies no-op when Mode == "bearer";
// extractors return "" on missing cookie so the middleware treats them as
// unauthenticated.
var SetAuthCookies = pkgauth.SetAuthCookies
var ClearAuthCookies = pkgauth.ClearAuthCookies
var ExtractAccessFromCookie = pkgauth.ExtractAccessFromCookie
var ExtractRefreshFromCookie = pkgauth.ExtractRefreshFromCookie

type CookieAttrs = pkgauth.CookieAttrs
`
	return os.WriteFile(filepath.Join(authDir, "reexport.go"), []byte(src), 0o644)
}
