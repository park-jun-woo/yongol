//ff:func feature=gen-gogin type=util control=sequence
//ff:what wrapAuthClaimsFields — auth.IssueToken/RefreshToken 시 Claims 래핑

package ssac

// wrapAuthClaimsFields wraps the mapped SSaC inputs in a
// `Claims: model.UserClaim{...}` literal when the target is
// auth.IssueToken / auth.RefreshToken. Phase001 UserClaimUnification —
// the claim payload type is the project-local model.UserClaim generated
// from manifest.backend.auth.claims; ssac/pkg/auth accepts any via the
// `Claims any` passthrough and JSON-marshals on issue.
func wrapAuthClaimsFields(pkgName, callFunc, fields string) string {
	if pkgName == "auth" && (callFunc == "IssueToken" || callFunc == "RefreshToken") {
		return "Claims: model.UserClaim{" + fields + "}"
	}
	return fields
}
