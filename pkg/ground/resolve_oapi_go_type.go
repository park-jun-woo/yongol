//ff:func feature=rule type=util control=selection topic=openapi
//ff:what resolveOAPIGoType — OpenAPI 스키마 ref → oapi-codegen Go 타입 (맥락·재귀·format-aware)

package ground

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// resolveOAPIGoType maps an OpenAPI schema reference to the Go type that
// oapi-codegen generates for it, recursively and context-aware:
//
//   - $ref            → the referenced type name (refName).
//   - array           → "[]" + resolveOAPIGoType(items, ctx)  (nested arrays
//     and array item formats are handled by the recursion — no per-shape
//     branch).
//   - string          → stringGoType(format, ctx)  (format×context table).
//   - integer/number  → context default, with int64/float overrides.
//   - boolean         → "bool".
//   - object          → "object" (response) / "" (param).
//
// Returns "" for nil refs, untyped schemas, unresolvable array items, or
// combinations a given context does not produce a Go type for. This unifies
// the former resolveOAPIParamGoType / resolveOAPIResponseGoType / responsePropType
// trio into one resolver so new shapes (nested arrays, array params) are
// covered without adding branches.
func resolveOAPIGoType(ref *openapi3.SchemaRef, ctx OAPIContext) string {
	if ref == nil {
		return ""
	}
	if ref.Ref != "" {
		return refName(ref.Ref)
	}
	s := ref.Value
	if s == nil || len(s.Type.Slice()) == 0 {
		return ""
	}
	switch s.Type.Slice()[0] {
	case "array":
		inner := resolveOAPIGoType(s.Items, ctx)
		if inner == "" {
			return ""
		}
		return "[]" + inner
	case "string":
		return stringGoType(s.Format, ctx)
	case "integer":
		if s.Format == "int64" {
			return "int64"
		}
		return ctxIntDefault(ctx)
	case "number":
		return ctxNumberType(s.Format, ctx)
	case "boolean":
		return "bool"
	case "object":
		if ctx == CtxParam {
			return ""
		}
		return "object"
	}
	return ""
}
