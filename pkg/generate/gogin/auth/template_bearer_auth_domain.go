//ff:type feature=gen-gogin type=generator
//ff:what bearerAuthStrictDomainTemplate — 도메인별 BearerAuthStrict<Title> 템플릿 (Bearer 헤더 인라인 추출, 전역 extractToken 우회)

package auth

// bearerAuthStrictDomainTemplate is the source for a per-domain
// internal/middleware/bearerauth_<ident>.go (Phase008 §3b/§3c). Placeholders:
//
//	__API_IMPORT__  domain api import path (<module>/internal/api_<ident>)
//	__MODULE__      module path (model import)
//	__API_PKG__     domain api package identifier (api_<ident>)
//	__TITLE__       PascalCase domain suffix (Public / Admin)
//
// Unlike the single-site bearerAuthStrictTemplate, the token is extracted by
// INLINING the Authorization `Bearer ` parse — the global extractToken()/
// authMode() switch is NOT consulted, so a mixed-mode binary (public=cookie +
// admin=bearer) reads the header unconditionally for this bearer domain.
const bearerAuthStrictDomainTemplate = `package middleware

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/park-jun-woo/ssac/pkg/auth"

	"__API_IMPORT__"
	"__MODULE__/internal/model"
)

// BearerAuthStrict__TITLE__ validates the session token for the __TITLE__
// domain per operation. Operations whose operationID is in publicOps bypass
// the check. The token source is the Authorization: Bearer header, inlined here
// (Phase008 §3c) so this domain never consults the global auth-mode switch.
func BearerAuthStrict__TITLE__(publicOps map[string]bool) __API_PKG__.StrictMiddlewareFunc {
	return func(handler __API_PKG__.StrictHandlerFunc, operationID string) __API_PKG__.StrictHandlerFunc {
		if publicOps[operationID] {
			return handler
		}
		return func(ctx *gin.Context, request interface{}) (interface{}, error) {
			token := ""
			if h := ctx.GetHeader("Authorization"); strings.HasPrefix(h, "Bearer ") {
				token = strings.TrimPrefix(h, "Bearer ")
			}
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
