//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildCallSetAuthCookiesLines — auth.RefreshToken 이후 SetAuthCookies 라인 (Phase020)

package ssac

import "fmt"

func (g *methodGen) buildCallSetAuthCookiesLines(varName string) []string {
	return []string{
		"// Phase020 — emit Set-Cookie for access+refresh. Runtime no-op when Mode==\"bearer\".",
		"if ginCtx, ok := ctx.(*gin.Context); ok {",
		fmt.Sprintf("\tauth.SetAuthCookies(ginCtx, %s.AccessToken, %s.RefreshToken)", g.AccessTokenVar, varName),
		"}",
	}
}
