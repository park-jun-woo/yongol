//ff:func feature=validate type=util control=iteration dimension=2 topic=ssac-structural
//ff:what xss60ResolveVarModel — 변수명에서 선행 @get/@post/@put/@delete result binding의 모델명 해결

package ssac

import parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"

// xss60ResolveVarModel resolves a variable name to its model type from
// preceding @get/@post/@put/@delete result bindings.
func xss60ResolveVarModel(varName string, fn parsessac.ServiceFunc) string {
	for _, seq := range fn.Sequences {
		switch seq.Type {
		case "get", "post", "put", "delete":
			if seq.Result != nil && seq.Result.Var == varName && seq.Result.Type != "" {
				return seq.Result.Type
			}
		}
	}
	return ""
}
