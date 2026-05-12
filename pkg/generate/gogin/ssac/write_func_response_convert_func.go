//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what writeFuncResponseConvertFunc — Func Response → api DTO converter 함수 본문 기록

package ssac

import (
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"

	"github.com/park-jun-woo/yongol/pkg/parser/funcspec"
)

// writeFuncResponseConvertFunc generates:
//
//	func convert<Name>(src <pkg>.<Type>) (*api.<Name>, error) {
//	    return &api.<Name>{
//	        OrgId: &src.OrgName,
//	        ...
//	    }, nil
//	}
//
// api side field names use pascalCase(jsonName) (oapi-codegen convention).
// Func side field names are looked up from FuncSpec.ResponseFields — matched
// by JSONName or caseconv.PascalToSnake(Field.Name) against the OpenAPI
// property's jsonName. Required fields use value assignment (src.Field);
// optional fields use pointer wrapping (&src.Field).
func writeFuncResponseConvertFunc(
	sb *strings.Builder,
	name string,
	schema *openapi3.Schema,
	info funcRespInfo,
	spec *funcspec.FuncSpec,
) {
	required := requiredSet(schema)

	propNames := make([]string, 0, len(schema.Properties))
	for k := range schema.Properties {
		propNames = append(propNames, k)
	}
	sort.Strings(propNames)

	// Build a lookup from jsonName → FuncSpec Field.Name for the Func side.
	funcFieldName := buildFuncFieldLookup(spec)

	sb.WriteString("func convert" + name + "(src " + info.PkgAlias + "." + name + ") (*api." + name + ", error) {\n")
	sb.WriteString("\treturn &api." + name + "{\n")
	for _, jsonName := range propNames {
		apiField := pascalCase(jsonName)
		srcField := funcFieldName[jsonName]
		if srcField == "" {
			// Fallback: use pascalCase(jsonName) — same as api side.
			srcField = apiField
		}
		if required[jsonName] {
			sb.WriteString("\t\t" + apiField + ": src." + srcField + ",\n")
		} else {
			sb.WriteString("\t\t" + apiField + ": &src." + srcField + ",\n")
		}
	}
	sb.WriteString("\t}, nil\n")
	sb.WriteString("}\n")
}
