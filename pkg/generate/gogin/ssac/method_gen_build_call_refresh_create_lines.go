//ff:func feature=gen-gogin type=util control=sequence
//ff:what buildCallRefreshCreateLines — auth.RefreshToken 이후 server.RefreshStore.Create 라인 (Phase004)

package ssac

import (
	"fmt"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func (g *methodGen) buildCallRefreshCreateLines(seq ssacparser.Sequence, varName string) []string {
	claimLit := "model.UserClaim{" + g.mapFields(seq.Inputs) + "}"
	return []string{
		fmt.Sprintf("if err := server.RefreshStore.Create(ctx, %s.RefreshToken, %s, %s.ExpiresAt); err != nil {", varName, claimLit, varName),
		fmt.Sprintf("\tslog.Error(\"handler: 5xx\", \"op\", %q, \"status\", 500, \"err\", err)", g.FuncName),
		fmt.Sprintf("\treturn api.%s500JSONResponse{Error: %q, Code: strPtr(%q)}, nil", g.FuncName, neutralMessage(500), neutralCode(500)),
		"}",
	}
}
