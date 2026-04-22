//ff:func feature=validate type=util control=sequence topic=openapi-ssac
//ff:what inferResponseValueType — @response field value 문자열에서 Go 타입 추론
package openapi_ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/rule"
)

// inferResponseValueType resolves a @response field expression to a Go type
// name via 3 paths:
//
//   1. Literal → inferLiteral (quoted string, numeric, bool, nil)
//   2. Bare variable ("user") → Types["SSaC.var.<funcName>.<var>"] as-is
//      (preserves slice/pointer/wrapper so []Webhook matches []Webhook)
//   3. Dotted (var.field) → type(var) via SSaC.var.*, strip wrapper/prefix,
//      then Types["Struct.<Type>.<field>"]
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
		varName := value[:dot]
		field := value[dot+1:]
		varType := g.Types["SSaC.var."+funcName+"."+varName]
		if varType == "" {
			return ""
		}
		// dotted field 조회 시 wrapper/slice/pointer/package prefix 모두 제거하여
		// Struct.<UnqualifiedTypeName>.<field> 로 normalize.
		return g.Types["Struct."+normalizeTypeName(varType)+"."+field]
	}
	return g.Types["SSaC.var."+funcName+"."+value]
}

