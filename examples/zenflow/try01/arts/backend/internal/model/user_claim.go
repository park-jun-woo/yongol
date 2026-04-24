//ff:type feature=model type=model
//ff:what UserClaim — JWT 인증 claim struct (manifest.backend.auth.claims 기반)
package model

// UserClaim carries the typed JWT claim fields for this project. It is both
// the payload passed to ssac/pkg/auth.IssueToken / RefreshToken via the
// Claims any passthrough (the shared runtime JSON-marshals the struct into
// jwt.MapClaims using these json tags) and the value BearerAuth middleware
// stores in the request ctx under the "currentUser" key for handlers to
// consume via ctx.Value("currentUser").(*model.UserClaim).
type UserClaim struct {
	Email string `json:"email"`
	ID int64 `json:"user_id"`
	OrgID int64 `json:"org_id"`
	Role string `json:"role"`
}
