//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-structural
//ff:what XOE-01 — ErrorResponse schema에 error/code가 required에 없으면 WARNING

package openapi

import (
	"fmt"
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoe01ErrorResponseRequired flags schemas whose name contains "Error" where
// the "error" or "code" property exists but is not listed in required.
// oapi-codegen generates *string for optional fields, which causes build
// failures when the generated code is consumed. Emitted as WARNING.
func xoe01ErrorResponseRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Components == nil || fs.OpenAPIDoc.Components.Schemas == nil {
		return nil
	}

	var diags []diagnostic.Diagnostic
	schemaNames := make([]string, 0, len(fs.OpenAPIDoc.Components.Schemas))
	for name := range fs.OpenAPIDoc.Components.Schemas {
		schemaNames = append(schemaNames, name)
	}
	sort.Strings(schemaNames)

	for _, name := range schemaNames {
		if !strings.Contains(name, "Error") {
			continue
		}
		ref := fs.OpenAPIDoc.Components.Schemas[name]
		if ref == nil || ref.Value == nil {
			continue
		}
		schema := ref.Value
		reqSet := make(map[string]bool, len(schema.Required))
		for _, r := range schema.Required {
			reqSet[r] = true
		}
		targets := []string{"error", "code"}
		for _, field := range targets {
			if _, hasProp := schema.Properties[field]; !hasProp {
				continue
			}
			if reqSet[field] {
				continue
			}
			line := fs.OpenAPILines.SchemaPropertyLine(name, field)
			if line == 0 {
				line = fs.OpenAPILines.SchemaLine(name)
			}
			diags = append(diags, diagnostic.Diagnostic{
				File:  "api/openapi.yaml",
				Line:  line,
				Phase: diagnostic.PhaseValidate,
				Level: diagnostic.LevelWarning,
				Message: fmt.Sprintf(
					"[XOE-01] ErrorResponse.%s가 required에 없으면 oapi-codegen이 *string으로 생성하여 codegen 산출물이 빌드 실패합니다. required에 추가하세요.",
					field),
				Advice: fmt.Sprintf("Add %q to the required list of schema %q", field, name),
			})
		}
	}
	return diags
}
