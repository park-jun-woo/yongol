//ff:func feature=gen-gogin type=util control=selection
//ff:what authStoreCookieEmit — RefreshRotate/Logout Set-Cookie emission block (Phase020)

package ssac

import "fmt"

// authStoreCookieEmit emits the Phase020 Set-Cookie block for auth
// RefreshRotate (fresh access+refresh cookies) and Logout (Max-Age=0 clear).
// Subscribe-path handlers emit nothing — cookies are an HTTP-only concern.
func (g *methodGen) authStoreCookieEmit(callFunc, varName string) []string {
	if g.IsSubscribe {
		return nil
	}
	switch callFunc {
	case "RefreshRotate":
		if varName == "_" {
			return nil
		}
		return []string{
			"// Phase020 — refresh rotation emits fresh Set-Cookie headers.",
			"if ginCtx, ok := ctx.(*gin.Context); ok {",
			fmt.Sprintf("\tauth.SetAuthCookies(ginCtx, %s.AccessToken, %s.RefreshToken)", varName, varName),
			"}",
		}
	case "Logout":
		return []string{
			"// Phase020 — Logout clears the session cookies (Max-Age=0).",
			"if ginCtx, ok := ctx.(*gin.Context); ok {",
			"\tauth.ClearAuthCookies(ginCtx)",
			"}",
		}
	default:
		return nil
	}
}
