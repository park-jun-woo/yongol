//ff:func feature=gen-gogin type=util control=sequence
//ff:what methodGen.buildResponse — @response 시퀀스 빌더 (convert 헬퍼 경유 + JSONB unmarshal 에러 전파)

package ssac

import (
	"fmt"

	ssacparser "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// buildResponse emits the final `return api.<Op><Status>JSONResponse(...)`
// line(s). When `@response <var>` targets a sqlc row variable, the row
// cannot be cast directly to the api DTO — sqlc and oapi-codegen diverge
// on acronym casing (OrgID vs OrgId), enum named types vs raw strings,
// and JSONB ↔ map[string]interface{} shape (BUG-003 / BUG-005). We route
// through convert<Model>(<var>) which now returns (*api.<Model>, error)
// so the handler can propagate unmarshal failures as 500 responses.
func (g *methodGen) buildResponse(seq ssacparser.Sequence) []string {
	// @response { workflow: updated } → field mapping via OpenAPI schema.
	if len(seq.Fields) > 0 {
		return g.buildFieldResponse(seq.Fields)
	}
	// @response target — direct variable (scalar or list).
	if seq.Target != "" {
		// @call result variables are Func Response types (user-authored,
		// OpenAPI-compatible) — emit direct cast without converter (BUG-050).
		if g.CallResultVars[seq.Target] {
			return []string{fmt.Sprintf("return api.%s%dJSONResponse(%s), nil",
				g.FuncName, g.SuccessStatus, seq.Target)}
		}
		if model := g.VarTypes[seq.Target]; model != "" {
			return []string{
				fmt.Sprintf("converted, err := convert%s(%s)", model, seq.Target),
				"if err != nil { return nil, err }",
				fmt.Sprintf("return api.%s%dJSONResponse(*converted), nil",
					g.FuncName, g.SuccessStatus),
			}
		}
		// No known model type (e.g. @response with a raw expression that
		// wasn't produced by a sqlc result binding). Keep the legacy
		// direct cast — this path was valid before and remains so for
		// scalar / wrapper responses.
		return []string{fmt.Sprintf("return api.%s%dJSONResponse(%s), nil",
			g.FuncName, g.SuccessStatus, seq.Target)}
	}
	// @response with no args — empty 2xx.
	return []string{fmt.Sprintf("return api.%s%dJSONResponse{}, nil",
		g.FuncName, g.SuccessStatus)}
}
