//ff:type feature=gen-hurl type=model
//ff:what authRole — scenarioCtx.authOpIDs 의 역할 태그 enum

package hurl

// authRole is the role tag stored in scenarioCtx.authOpIDs for O(1)
// "is this op an auth op?" lookup during downstream phase building.
type authRole string

const (
	authRoleSignup authRole = "signup"
	authRoleLogin  authRole = "login"
)
