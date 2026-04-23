//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what writeConvertFunc — convert<Name>(row db.X) *api.X 함수 본문 기록

package ssac

import (
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// writeConvertFunc generates: func convertWorkflow(row db.Workflow) *api.Workflow
//
// Only optional OpenAPI properties are wrapped with ptrOf — oapi-codegen
// emits optional fields as *T, required fields as T, so wrapping a
// required field produces a *T that won't assign to the T struct slot
// (BUG-006).
func writeConvertFunc(sb *strings.Builder, name string, schema *openapi3.Schema) {
	required := requiredSet(schema)

	sb.WriteString("func convert" + name + "(row db." + name + ") *api." + name + " {\n")
	sb.WriteString("\treturn &api." + name + "{\n")

	propNames := make([]string, 0, len(schema.Properties))
	for k := range schema.Properties {
		propNames = append(propNames, k)
	}
	sort.Strings(propNames)

	for _, jsonName := range propNames {
		apiField := pascalCase(jsonName)    // oapi-codegen: "org_id" → "OrgId"
		dbField := sqlcPascalCase(jsonName) // sqlc: "org_id" → "OrgID" (Go acronyms)
		if required[jsonName] {
			// Required → value type on api side. Assign directly.
			sb.WriteString("\t\t" + apiField + ": row." + dbField + ",\n")
		} else {
			// Optional → *T on api side. Wrap.
			sb.WriteString("\t\t" + apiField + ": ptrOf(row." + dbField + "),\n")
		}
	}

	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
}

// requiredSet flattens schema.Required into a lookup set. Returns an
// empty (non-nil) map when schema is nil or has no required list so
// callers can index without a nil check.
func requiredSet(schema *openapi3.Schema) map[string]bool {
	out := make(map[string]bool, len(schema.Required))
	if schema == nil {
		return out
	}
	for _, r := range schema.Required {
		out[r] = true
	}
	return out
}
