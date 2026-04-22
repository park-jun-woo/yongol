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
//  2. bare variable via Types[SSaC.var.*]
//
// Field access (var.Field) returns "" (deferred).
func resolveInputType(g *rule.Ground, funcName, value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	if t := inferLiteralType(value); t != "" {
		return t
	}
	if strings.ContainsAny(value, ".\"'") {
		return ""
	}
	return g.Types["SSaC.var."+funcName+"."+value]
}
