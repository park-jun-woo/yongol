//ff:func feature=gen-gogin type=util control=sequence
//ff:what methodGen.resolveErrCode — 요청한 HTTP status를 그대로 반환

package ssac

// resolveErrCode returns the requested status as-is.
// Validate (XOS-21) guarantees every ErrStatus used in SSaC is defined in OpenAPI.
// No fallback needed.
func (g *methodGen) resolveErrCode(wanted int) int {
	return wanted
}
