//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildCallRefreshCreateLines — auth.RefreshToken 이후 auth.CreateRefresh 라인 (Phase004/Phase002 ssac/purify)

package ssac

import (
	"fmt"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// buildCallRefreshCreateLines emits the persistence call that follows a
// successful auth.RefreshToken in SSaC sequences (e.g. Login → refresh token
// issuance → row insert).
//
// Phase002 (ssac/purify) — server.RefreshStore is gone. ssac/pkg/auth now
// exposes auth.CreateRefresh(ctx, token, claims, expiresAt) which reads the
// package-level singleton installed via auth.Init(...). Handler code never
// references the Server struct for this call.
func (g *methodGen) buildCallRefreshCreateLines(seq ssacparser.Sequence, varName string) []string {
	claimLit := "model.UserClaim{" + g.mapFields(seq.Inputs) + "}"
	return []string{
		fmt.Sprintf("if err := auth.CreateRefresh(ctx, %s.RefreshToken, %s, %s.ExpiresAt); err != nil {", varName, claimLit, varName),
		fmt.Sprintf("\tslog.Error(\"handler: 5xx\", \"op\", %q, \"status\", 500, \"err\", err)", g.FuncName),
		fmt.Sprintf("\treturn api.%s500JSONResponse{Error: %q, Code: strPtr(%q)}, nil", g.FuncName, neutralMessage(500), neutralCode(500)),
		"}",
	}
}
