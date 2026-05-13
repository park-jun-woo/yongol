package auth

// authModeTemplate is the source for internal/middleware/auth_mode.go.
// %s is the default auth mode literal.
var authModeTemplate = `package middleware

import (
	"os"
	"strings"
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
`

// extractTokenTemplate is the source for internal/middleware/extract_token.go.
var extractTokenTemplate = `package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/park-jun-woo/ssac/pkg/auth"
)

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
`

// bearerAuthStrictTemplate is the source for internal/middleware/bearerauth.go.
// %s verbs: 1=modulePath (api), 2=modulePath (model).
var bearerAuthStrictTemplate = `package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/park-jun-woo/ssac/pkg/auth"

	"%s/internal/api"
	"%s/internal/model"
)

// BearerAuthStrict returns a StrictMiddlewareFunc that validates the
// session token per operation. Operations whose operationID is in
// publicOps bypass the check (e.g. Login, Register).
//
// Phase020 — the token source depends on auth mode (see extractToken).
// Phase003 — the signing secret is sourced inside auth.VerifyToken via
// auth.Configure(SecretEnv) → os.Getenv; no secret parameter here.
// Phase001 UserClaimUnification — the JSON round-trip unmarshals directly
// into model.UserClaim and stores &claim in the ctx; the previous field-
// by-field copy into a separate ctx type is gone.
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
			raw, err := json.Marshal(out.Claims)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				return nil, nil
			}
			var claim model.UserClaim
			if err := json.Unmarshal(raw, &claim); err != nil {
				ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				return nil, nil
			}
			ctx.Set("currentUser", &claim)
			return handler(ctx, request)
		}
	}
}
`
