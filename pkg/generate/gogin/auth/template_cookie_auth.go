//ff:type feature=gen-gogin type=generator
//ff:what cookieAuthStrictDomainTemplate — 도메인별 CookieAuthStrict<Title> 템플릿 (쿠키 토큰 직접 추출, 전역 extractToken 우회)

package auth

// cookieAuthStrictDomainTemplate is the source for a per-domain
// internal/middleware/cookieauth_<ident>.go (Phase008 §3a/§3c). Placeholders
// match bearerAuthStrictDomainTemplate (__API_IMPORT__ / __MODULE__ /
// __API_PKG__ / __TITLE__). The token is read straight from the session cookie
// via auth.ExtractAccessFromCookie(ctx) — the bearer template's extractToken()/
// authMode() path is intentionally NOT used, so a cookie domain always reads
// the cookie regardless of the global BACKEND_AUTH_MODE.
const cookieAuthStrictDomainTemplate = `package middleware

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/park-jun-woo/ssac/pkg/auth"

	"__API_IMPORT__"
	"__MODULE__/internal/model"
)

// CookieAuthStrict__TITLE__ validates the session token for the __TITLE__
// domain per operation. Operations whose operationID is in publicOps bypass the
// check. The token is extracted directly from the session cookie (Phase008
// §3c) so this domain never consults the global auth-mode switch.
func CookieAuthStrict__TITLE__(publicOps map[string]bool) __API_PKG__.StrictMiddlewareFunc {
	return func(handler __API_PKG__.StrictHandlerFunc, operationID string) __API_PKG__.StrictHandlerFunc {
		if publicOps[operationID] {
			return handler
		}
		return func(ctx *gin.Context, request interface{}) (interface{}, error) {
			token := auth.ExtractAccessFromCookie(ctx)
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
