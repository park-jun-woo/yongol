//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what generateVerifyMiddleware — internal/middleware/bearerauth.go 생성 (mode 기반 토큰 추출)
package auth

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/gogin/ffannot"
)

// generateBearerAuth writes internal/middleware/bearerauth.go — a
// StrictMiddlewareFunc that validates session tokens per operation. The
// middleware consults a publicOps map (passed in by main.go from
// collectPublicOps) to bypass auth for endpoints marked `security: []` in
// OpenAPI.
//
// Phase020 — the middleware now branches on the manifest-resolved auth mode
// at request time:
//
//   - bearer: read the token from the Authorization header only.
//   - cookie: read the token from __Host-access_token (or the configured
//     CookieAttrs.AccessName) only.
//   - hybrid: prefer Authorization; fall back to the cookie when the
//     header is absent. Mobile clients emit the header; browsers emit the
//     cookie; both succeed against the same backend.
//
// The mode is resolved from os.Getenv("BACKEND_AUTH_MODE") at request time
// with a generator-baked default (the manifest-resolved value). This mirrors
// the three-tier pattern used elsewhere (manifest default < env override <
// explicit header).
//
// The filename stays bearerauth.go (rename deferred) so this Phase doesn't
// sweep unrelated imports in main.go; the file's //ff:what now reflects the
// broader responsibility.
func generateBearerAuth(artifactsDir, modulePath string, fields []ClaimField, defaultMode string) error {
	mwDir := filepath.Join(artifactsDir, "backend", "internal", "middleware")
	if err := os.MkdirAll(mwDir, 0o755); err != nil {
		return err
	}

	var assignments []string
	for _, f := range fields {
		assignments = append(assignments, fmt.Sprintf("\t\t\t\t%s: claim.%s,", f.Name, f.Name))
	}

	header := ffannot.EmitAnnotationBlock(ffannot.Block{
		Func: ffannot.FuncAnnot{Feature: "middleware", Type: "middleware", Control: "sequence", Topic: "auth-check"},
		What: "BearerAuthStrict — oapi-codegen per-op 세션 토큰 검증 미들웨어 (mode 기반 분기, ssac/pkg/auth 기반)",
	})
	src := header + fmt.Sprintf(`package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"%s/internal/api"
	"%s/internal/auth"
	"%s/internal/model"
)

// authMode returns the effective auth transport mode. BACKEND_AUTH_MODE
// env overrides the manifest-resolved default; values outside the closed
// set {bearer, cookie, hybrid} fall back to the default.
func authMode() string {
	v := strings.ToLower(strings.TrimSpace(os.Getenv("BACKEND_AUTH_MODE")))
	switch v {
	case "bearer", "cookie", "hybrid":
		return v
	}
	return %q
}

// extractToken reads the session JWT according to the current auth mode.
// Return "" signals "no token present" — the middleware converts that to
// a 401 without distinguishing missing vs malformed.
func extractToken(ctx *gin.Context) string {
	mode := authMode()
	header := ctx.GetHeader("Authorization")
	var headerToken string
	if strings.HasPrefix(header, "Bearer ") {
		headerToken = strings.TrimPrefix(header, "Bearer ")
	}
	switch mode {
	case "bearer":
		return headerToken
	case "cookie":
		return auth.ExtractAccessFromCookie(ctx)
	case "hybrid":
		if headerToken != "" {
			return headerToken
		}
		return auth.ExtractAccessFromCookie(ctx)
	}
	return headerToken
}

// BearerAuthStrict returns a StrictMiddlewareFunc that validates the
// session token per operation. Operations whose operationID is in
// publicOps bypass the check (e.g. Login, Register).
//
// Phase020 — the token source depends on auth mode (see extractToken).
// Phase003 — the signing secret is sourced inside auth.VerifyToken via
// auth.Configure(SecretEnv) → os.Getenv; no secret parameter here.
func BearerAuthStrict(publicOps map[string]bool) api.StrictMiddlewareFunc {
	return func(handler api.StrictHandlerFunc, operationID string) api.StrictHandlerFunc {
		if publicOps[operationID] {
			return handler
		}
		return func(ctx *gin.Context, request interface{}) (interface{}, error) {
			token := extractToken(ctx)
			if token == "" {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				return nil, nil
			}
			out, err := auth.VerifyToken(auth.VerifyTokenRequest{Token: token})
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				return nil, nil
			}
			// MapClaims → typed Claim via JSON round-trip. Keeps this file
			// agnostic to the concrete claim field list; the Claim struct
			// (generated from manifest.backend.auth.claims) carries the
			// json tags that match the JWT claim keys.
			raw, err := json.Marshal(out.Claims)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				return nil, nil
			}
			var claim auth.Claim
			if err := json.Unmarshal(raw, &claim); err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				return nil, nil
			}
			ctx.Set("currentUser", &model.CurrentUser{
%s
			})
			return handler(ctx, request)
		}
	}
}
`, modulePath, modulePath, modulePath, defaultMode, strings.Join(assignments, "\n"))

	return os.WriteFile(filepath.Join(mwDir, "bearerauth.go"), []byte(src), 0o644)
}
