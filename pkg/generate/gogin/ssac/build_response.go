//ff:func feature=gen-gogin type=util control=sequence
//ff:what methodGen.buildResponse — @response 시퀀스 빌더 (OpenAPI 200 response 매핑)

package ssac

import (
	"fmt"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func (g *methodGen) buildResponse(seq ssacparser.Sequence) []string {
	// @response { workflow: updated } → field mapping via OpenAPI schema
	if len(seq.Fields) > 0 {
		return g.buildFieldResponse(seq.Fields)
	}
	// @response target — direct variable (e.g. list result). g.SuccessStatus
	// was derived from the operation's HTTP method + declared 2xx responses
	// at extract time so POST handlers emit 201, DELETE handlers emit 204,
	// etc. instead of the previous hardcoded 200 (BUG-004).
	if seq.Target != "" {
		return []string{fmt.Sprintf("return api.%s%dJSONResponse(%s), nil",
			g.FuncName, g.SuccessStatus, seq.Target)}
	}
	// @response with no args — empty 2xx
	return []string{fmt.Sprintf("return api.%s%dJSONResponse{}, nil",
		g.FuncName, g.SuccessStatus)}
}
