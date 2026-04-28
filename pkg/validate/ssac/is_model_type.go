//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what isModelType — typeName이 DDL Model 또는 등록된 struct (Func Response 등) 인지 판정

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// isModelType reports whether typeName is a DDL Model (registered in
// Ground.Lookup["SymbolTable.model"]) or any struct registered in Ground.Types
// under "Struct.<typeName>.*" (e.g. a Func Response struct emitted by
// registerFuncSpec, or a DDL row struct emitted by populateSSaCSymbols).
func isModelType(g *rule.Ground, typeName string) bool {
	if g == nil {
		return false
	}
	if models, ok := g.Lookup["SymbolTable.model"]; ok {
		if models[typeName] {
			return true
		}
	}
	prefix := "Struct." + typeName + "."
	for k := range g.Types {
		if strings.HasPrefix(k, prefix) {
			return true
		}
	}
	return false
}
