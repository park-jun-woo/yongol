//ff:func feature=validate type=util control=sequence topic=openapi-ssac
//ff:what inferDottedFieldType — var.field 표현식을 SSaC.var → DDL.apifield/Struct 경로로 Go 타입 해석
package openapi_ssac

import "github.com/park-jun-woo/yongol/pkg/rule"

// inferDottedFieldType resolves a dotted @response field (var.field) to a Go
// type. It looks up the variable's type, normalises it to an unqualified model
// name, then prefers the DDL api-surface type over the coarse Struct.*
// projection.
//
// When looking up a dotted field, all wrapper/slice/pointer/package prefixes
// are stripped and normalised to <UnqualifiedTypeName>.<field>.
//
// The DDL.apifield key (e.g. a UUID column's openapi_types.UUID) is preferred
// over the GoTypeOf projection in Struct.* (which collapses UUID→string). It
// falls back to Struct.* for func-result structs and non-DDL row types, where
// no apifield key exists. Both keys share the same <Model>.<field> token space
// (populate_ddl.go and populate_ssac_symbols.go use the same casing functions —
// pinned by TestPopulateDDL_ApifieldStructKeyParity).
func inferDottedFieldType(g *rule.Ground, funcName, varName, field string) string {
	varType := g.Types["SSaC.var."+funcName+"."+varName]
	if varType == "" {
		return ""
	}
	model := normalizeTypeName(varType)
	if apiType := g.Types["DDL.apifield."+model+"."+field]; apiType != "" {
		return apiType
	}
	return g.Types["Struct."+model+"."+field]
}
