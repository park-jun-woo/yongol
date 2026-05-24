//ff:func feature=gen-gogin type=util control=selection
//ff:what methodGen.buildResponse — @response 시퀀스 빌더 (convert 헬퍼 경유 + JSONB unmarshal 에러 전파 + pgtype 변환)

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
//
// PhaseG02 — returns (lines, imports) instead of lines-only so that field
// response pgtype conversions can propagate their pgtypex import paths to
// the method file writer.
func (g *methodGen) buildResponse(seq ssacparser.Sequence) ([]string, []string) {
	// @response { workflow: updated } → field mapping via OpenAPI schema.
	if len(seq.Fields) > 0 {
		return g.buildFieldResponse(seq.Fields)
	}
	// @response target — direct variable (scalar or list).
	if seq.Target != "" {
		if model := g.VarTypes[seq.Target]; model != "" {
			return g.buildResponseConvert(model, seq.Target), nil
		}
		// No known model type (e.g. @response with a raw expression that
		// wasn't produced by a sqlc result binding). Keep the legacy
		// direct cast — this path was valid before and remains so for
		// scalar / wrapper responses.
		return []string{fmt.Sprintf("return api.%s%dJSONResponse(%s), nil",
			g.FuncName, g.SuccessStatus, seq.Target)}, nil
	}
	// @response with no args — empty 2xx.
	suffix := "JSONResponse"
	if g.SuccessStatus == 204 {
		suffix = "Response"
	}
	return []string{fmt.Sprintf("return api.%s%d%s{}, nil",
		g.FuncName, g.SuccessStatus, suffix)}, nil
}
