//ff:func feature=gen-gogin type=util control=sequence
//ff:what wrapAuthClaimsFields — auth.IssueToken/RefreshToken 시 Claims 래핑

package ssac

func wrapAuthClaimsFields(pkgName, callFunc, fields string) string {
	if pkgName == "auth" && (callFunc == "IssueToken" || callFunc == "RefreshToken") {
		return "Claims: auth.Claim{" + fields + "}"
	}
	return fields
}
