//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what methodGen.sqlcArgsSingle — 입력이 1개일 때 sqlc 인자 문자열 생성 (JSONB hoist 포함)
package ssac

func (g *methodGen) sqlcArgsSingle(inputs map[string]string) (preamble []string, args string, imports []string) {
	for k, v := range inputs {
		raw, pre, needsJSON := g.maybeMarshalJSONB(k, v)
		if needsJSON {
			imports = append(imports, `"encoding/json"`)
			return pre, "ctx, " + raw, imports
		}
		return nil, "ctx, " + g.mapValue(v), nil
	}
	return nil, "ctx", nil
}
