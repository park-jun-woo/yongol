//ff:func feature=validate type=util control=sequence topic=openapi-ssac
//ff:what inferResponseValueType — infers a Go type from a @response field value string
package openapi_ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// inferResponseValueType resolves a @response field expression to a Go type
// name via 3 paths:
//
//  1. Literal → inferLiteral (quoted string, numeric, bool, nil)
//  2. Bare variable ("user") → Types["SSaC.var.<funcName>.<var>"] as-is
//     (preserves slice/pointer/wrapper so []Webhook matches []Webhook)
//  3. Dotted (var.field) → type(var) via SSaC.var.*, strip wrapper/prefix,
//     then Types["Struct.<Type>.<field>"]
//
// Returns "" when unresolvable (downstream skips comparison).
func inferResponseValueType(g *rule.Ground, funcName, value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return ""
	}
	if t := inferLiteral(value); t != "" {
		return t
	}
	if dot := strings.IndexByte(value, '.'); dot > 0 {
		return inferDottedFieldType(g, funcName, value[:dot], value[dot+1:])
	}
	return g.Types["SSaC.var."+funcName+"."+value]
}
