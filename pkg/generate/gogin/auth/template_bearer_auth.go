package auth

// bearerAuthTemplate is the raw Go source emitted into
// internal/middleware/bearerauth.go. The three %s verbs bind, in order:
//
//  1. modulePath for the api import
//  2. modulePath for the model import
//  3. the default auth mode literal returned when BACKEND_AUTH_MODE is unset
//     or invalid
//
// Kept as a var so generateBearerAuth stays under the Q3 sequence line budget.
var bearerAuthTemplate = `package middleware

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/park-jun-woo/ssac/pkg/auth"

	"%s/internal/api"
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
			// MapClaims → typed UserClaim via JSON round-trip. Keeps this
			// file agnostic to the concrete claim field list; the
			// UserClaim struct (generated from manifest.backend.auth.claims)
			// carries the json tags that match the JWT claim keys.
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
