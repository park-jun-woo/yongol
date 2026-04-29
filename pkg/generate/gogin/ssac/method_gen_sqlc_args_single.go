//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what methodGen.sqlcArgsSingle — 입력이 1개일 때 sqlc 인자 문자열 생성 (JSONB hoist + 리터럴 wrap)
package ssac

func (g *methodGen) sqlcArgsSingle(inputs map[string]string) (preamble []string, args string, imports []string) {
	for k, v := range inputs {
		raw, pre, needsJSON := g.maybeMarshalJSONB(k, v)
		if needsJSON {
			imports = append(imports, `"encoding/json"`)
			return pre, "ctx, " + raw, imports
		}
		rendered := g.mapValue(v)
		// BUG-037 #1 — string literal at a JSONB column requires
		// []byte(...) wrap so sqlc params accept it.
		rendered = g.wrapJSONBLiteral(k, rendered)
		return nil, "ctx, " + rendered, nil
	}
	return nil, "ctx", nil
}
