//ff:func feature=rule type=loader control=iteration dimension=1
//ff:what registerFuncSpec — 하나의 FuncSpec 를 Ground Types/Schemas 에 등록

package ground

import (
	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// registerFuncSpec populates Ground Types / Schemas for a single FuncSpec.
// camelCase @func annotation name is preserved as the canonical key.
func registerFuncSpec(g *rule.Ground, sp *funcspec.FuncSpec) {
	funcPascal := capitalizeFirst(sp.Name)
	reqTypeName := funcPascal + "Request"
	respTypeName := funcPascal + "Response"

	var reqFields []string
	for _, f := range sp.RequestFields {
		reqFields = append(reqFields, f.Name)
		g.Types["Func.request."+sp.Name+"."+f.Name] = f.Type
		g.Types["Struct."+reqTypeName+"."+f.Name] = f.Type
	}
	g.Schemas["Func.request."+sp.Name] = reqFields

	for _, f := range sp.ResponseFields {
		g.Types["Struct."+respTypeName+"."+f.Name] = f.Type
	}
}
