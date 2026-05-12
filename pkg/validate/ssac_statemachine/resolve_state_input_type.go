//ff:func feature=validate type=util control=sequence topic=ssac-statemachine
//ff:what resolveStateInputType — @state input value 표현식의 Go 타입을 Ground.Types 에서 해석

package ssac_statemachine

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/rule"
)


// resolveStateInputType resolves the Go type of a @state input value
// expression using Ground.Types. Duplicates the var.Field logic from
// ssac_func.resolveInputType to avoid cross-package dependency.
func resolveStateInputType(g *rule.Ground, funcName, value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	if t := inferStateLiteralType(value); t != "" {
		return t
	}
	if strings.HasPrefix(value, "currentUser.") {
		field := value[len("currentUser."):]
		return g.Types["Manifest.claim."+field]
	}
	if dot := strings.IndexByte(value, '.'); dot > 0 {
		return resolveStateVarField(g, funcName, value[:dot], value[dot+1:])
	}
	if strings.ContainsAny(value, ".\"'") {
		return ""
	}
	return g.Types["SSaC.var."+funcName+"."+value]
}
