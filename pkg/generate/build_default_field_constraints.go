//ff:func feature=generate type=util control=iteration dimension=1
//ff:what OpenAPI 스키마의 properties에서 기본 타입의 FieldConstraint 맵을 생성한다

package generate

import (
	"github.com/getkin/kin-openapi/openapi3"

	oapiparser "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// buildDefaultFieldConstraints creates a FieldConstraint map from an OpenAPI
// schema's properties, using only the type and required information. This is
// used as a fallback when the operation was not captured by the primary
// ExtractRequestConstraints pipeline.
func buildDefaultFieldConstraints(schema *openapi3.Schema) map[string]oapiparser.FieldConstraint {
	if schema == nil || len(schema.Properties) == 0 {
		return nil
	}
	reqSet := make(map[string]bool, len(schema.Required))
	for _, r := range schema.Required {
		reqSet[r] = true
	}
	fields := make(map[string]oapiparser.FieldConstraint, len(schema.Properties))
	for name, ref := range schema.Properties {
		if ref.Value == nil {
			continue
		}
		var typeName string
		if types := ref.Value.Type.Slice(); len(types) > 0 {
			typeName = types[0]
		}
		fields[name] = oapiparser.FieldConstraint{
			Type:     typeName,
			Format:   ref.Value.Format,
			Required: reqSet[name],
		}
	}
	return fields
}
