//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what methodGen.sqlcArgsSingle — 입력이 1개일 때 sqlc 인자 문자열 생성 (JSONB hoist + InsertExpr bridge + 리터럴 wrap)
package ssac

import "strings"

func (g *methodGen) sqlcArgsSingle(inputs map[string]string) (preamble []string, args string, imports []string) {
	for k, v := range inputs {
		raw, pre, needsJSON := g.maybeMarshalJSONB(k, v)
		if needsJSON {
			imports = append(imports, `"encoding/json"`)
			return pre, "ctx, " + raw, imports
		}
		rendered := g.mapValue(v)
		rendered = g.wrapJSONBLiteral(k, rendered)
		alreadyPgtype := !strings.HasPrefix(v, "request.") && !strings.HasPrefix(v, `"`)
		rendered, extraImports := g.wrapInsertExpr(k, rendered, alreadyPgtype)
		imports = append(imports, extraImports...)
		return nil, "ctx, " + rendered, imports
	}
	return nil, "ctx", nil
}
