//ff:func feature=validate type=util control=sequence topic=ssac-statemachine
//ff:what resolveStateVarField — var.Field 표현식을 SSaC.var -> DDL.field 체인으로 Go 타입 해석

package ssac_statemachine

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// resolveStateVarField resolves a var.Field expression via SSaC.var -> DDL.field.
func resolveStateVarField(g *rule.Ground, funcName, varName, fieldName string) string {
	modelType := g.Types["SSaC.var."+funcName+"."+varName]
	if modelType == "" {
		return ""
	}
	modelType = strings.TrimPrefix(strings.TrimPrefix(modelType, "[]"), "*")
	return g.Types["DDL.field."+modelType+"."+fieldName]
}
