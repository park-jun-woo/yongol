//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what writeConvertFunc — convert<Name>(row db.X) (*api.X, error) 함수 본문 기록

package ssac

import (
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// writeConvertFunc generates: func convertWorkflow(row db.Workflow) (*api.Workflow, error)
//
// Two layer-crossings are handled:
//
//  1. Required vs optional fields (BUG-006) — oapi-codegen emits optional
//     properties as *T and required as T. Wrapping a required T with
//     ptrOf yields *T and fails to assign, so required fields go straight
//     through as values while optional ones get ptrOf.
//
//  2. JSONB ↔ map[string]interface{} boundary (BUG-005) — when an
//     OpenAPI property is `type: object, additionalProperties: true`,
//     oapi-codegen emits a map[string]interface{} while sqlc emits
//     json.RawMessage for the JSONB column. Direct assignment is a
//     Go-level type error. writeConvertFunc declares a local variable of
//     the map type and `json.Unmarshal`s the raw bytes into it before
//     the struct literal. Any unmarshal failure propagates as the
//     convert's second return value so callers surface the HTTP 500 /
//     trigger tx rollback at the transport layer.
//
// The convert signature always returns an error — even for schemas with
// no JSONB columns — so every caller uses a single pattern. The extra
// nil return is a no-op at runtime.
func writeConvertFunc(sb *strings.Builder, name string, schema *openapi3.Schema) {
	required := requiredSet(schema)

	propNames := make([]string, 0, len(schema.Properties))
	for k := range schema.Properties {
		propNames = append(propNames, k)
	}
	sort.Strings(propNames)

	jsonbs := collectJSONBFieldAliases(schema, propNames)

	sb.WriteString("func convert" + name + "(row db." + name + ") (*api." + name + ", error) {\n")
	writeJSONBUnmarshalScaffolding(sb, jsonbs)
	sb.WriteString("\treturn &api." + name + "{\n")
	for _, jsonName := range propNames {
		apiField := pascalCase(jsonName)
		dbField := sqlcPascalCase(jsonName)
		apiType := apiCastFor(name, jsonName, schema.Properties[jsonName])
		rhs := pickConvertRHS(jsonName, apiField, dbField, required[jsonName], jsonbs, apiType)
		sb.WriteString("\t\t" + apiField + ": " + rhs + ",\n")
	}
	sb.WriteString("\t}, nil\n")
	sb.WriteString("}\n")
}
