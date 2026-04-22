//ff:func feature=gen-gogin type=generator control=iteration dimension=1
//ff:what writeConvertFunc — convert<Name>(row db.X) *api.X 함수 본문 기록

package ssac

import (
	"sort"
	"strings"

	"github.com/getkin/kin-openapi/openapi3"
)

// writeConvertFunc generates: func convertWorkflow(row db.Workflow) *api.Workflow
func writeConvertFunc(sb *strings.Builder, name string, schema *openapi3.Schema) {
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
		sb.WriteString("\t\t" + apiField + ": ptrOf(row." + dbField + "),\n")
	}

	sb.WriteString("\t}\n")
	sb.WriteString("}\n")
}
