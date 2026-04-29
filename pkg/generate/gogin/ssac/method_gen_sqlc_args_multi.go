//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what methodGen.sqlcArgsMulti — 입력이 2개 이상일 때 sqlc Params 구조체 리터럴 생성 (JSONB 리터럴 wrap 포함)
package ssac

import (
	"sort"
	"strings"
)

// Validate (XQS-14/16) guarantees key matches sqlc field name. Use as-is.
func (g *methodGen) sqlcArgsMulti(method string, inputs map[string]string) (preamble []string, args string, imports []string) {
	var fields []string
	for k, v := range inputs {
		raw, pre, needsJSON := g.maybeMarshalJSONB(k, v)
		if needsJSON {
			preamble = append(preamble, pre...)
			imports = append(imports, `"encoding/json"`)
			fields = append(fields, k+": "+raw)
			continue
		}
		rendered := g.mapValue(v)
		// BUG-037 #1 — string literal at a JSONB column requires
		// []byte(...) wrap so sqlc params accept it.
		rendered = g.wrapJSONBLiteral(k, rendered)
		fields = append(fields, k+": "+rendered)
	}
	sort.Strings(fields)
	return preamble, "ctx, db." + method + "Params{" + strings.Join(fields, ", ") + "}", imports
}
