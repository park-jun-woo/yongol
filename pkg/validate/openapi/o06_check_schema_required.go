//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-structural
//ff:what o06CheckSchemaRequired — 단일 스키마의 required 항목이 properties 에 없으면 O-6 ERROR

package openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// o06CheckSchemaRequired emits one [O-6] ERROR per entry in the schema's
// required list that is not declared in its properties (a dangling required).
// Line resolution is delegated to o06RequiredLine.
func o06CheckSchemaRequired(fs *yongol.Fullstack, entry o06SchemaEntry) []diagnostic.Diagnostic {
	s := entry.schema
	if s == nil || len(s.Required) == 0 {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, name := range s.Required {
		if _, ok := s.Properties[name]; ok {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    "api/openapi.yaml",
			Line:    o06RequiredLine(fs, entry.schemaName, name),
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[O-6] required 항목 %q이(가) 같은 스키마의 properties에 선언돼 있지 않습니다(dangling required).", name),
			Advice:  fmt.Sprintf("Declare %q under properties, or remove it from the required list.", name),
		})
	}
	return diags
}
