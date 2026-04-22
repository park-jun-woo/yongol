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
	// @response target — direct variable (e.g. list result)
	if seq.Target != "" {
		return []string{fmt.Sprintf("return api.%s200JSONResponse(%s), nil", g.FuncName, seq.Target)}
	}
	// @response with no args — empty 200
	return []string{fmt.Sprintf("return api.%s200JSONResponse{}, nil", g.FuncName)}
}
