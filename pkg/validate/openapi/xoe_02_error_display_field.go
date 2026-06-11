//ff:func feature=validate type=rule control=iteration dimension=2 topic=openapi-structural
//ff:what XOE-02 — codegen이 에러 표시에 쓰는 필드가 ErrorResponse schema에 string 속성으로 실재하지 않으면 WARNING

package openapi

import (
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoe02ErrorDisplayField flags schemas whose name contains "Error" where the
// field a generated mutation onError handler reads (ExtractErrorDisplayField:
// "error" first, "message" second) is not present as a string property. When
// neither field exists, or exists with a non-string type, the generated
// frontend can only surface action failures through the String(err) fallback
// ("[object Object]" for a plain ErrorResponse, BUG-125). This is a display
// quality regression rather than a runtime break, so it is emitted as WARNING
// (matching XOE-01's severity) to avoid breaking projects that already pass.
func xoe02ErrorDisplayField(fs *yongol.Fullstack) []diagnostic.Diagnostic {
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
		if d := xoe02SchemaDiag(name, ref.Value, fs.OpenAPILines); d != nil {
			diags = append(diags, *d)
		}
	}
	return diags
}
