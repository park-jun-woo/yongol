//ff:func feature=validate type=util control=sequence topic=func-check
//ff:what resolveInputType — resolve the Go type of a @call input value expression (literal or SSaC.var lookup)

package ssac_func

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// resolveInputType returns the Go type of an @call input value expression.
// Resolution order:
//  1. inferLiteralType (quoted string, numeric, bool, nil)
//  2. currentUser.<Field> → Manifest.claim.<Field>
//  3. var.Field → SSaC.var → DDL.field (Phase009)
//  4. bare variable via Types[SSaC.var.*]
func resolveInputType(g *rule.Ground, funcName, value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	if t := inferLiteralType(value); t != "" {
		return t
	}
	if strings.HasPrefix(value, "currentUser.") {
		field := value[len("currentUser."):]
		return g.Types["Manifest.claim."+field]
	}
	// var.Field: resolve model type then DDL column Go type
	if dot := strings.IndexByte(value, '.'); dot > 0 {
		varName := value[:dot]
		fieldName := value[dot+1:]
		modelType := g.Types["SSaC.var."+funcName+"."+varName]
		if modelType != "" {
			// strip slice/pointer prefix: "[]Workflow" → "Workflow", "*User" → "User"
			modelType = strings.TrimPrefix(strings.TrimPrefix(modelType, "[]"), "*")
			return g.Types["DDL.field."+modelType+"."+fieldName]
		}
	}
	if strings.ContainsAny(value, ".\"'") {
		return ""
	}
	return g.Types["SSaC.var."+funcName+"."+value]
}
