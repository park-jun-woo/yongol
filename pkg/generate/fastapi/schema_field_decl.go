//ff:func feature=gen-fastapi type=util control=sequence
//ff:what schemaFieldDecl — 단일 body 필드 → Pydantic 모델 필드 선언 줄 (required/optional)

package fastapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// schemaFieldDecl returns the Pydantic model field declaration line for a body
// field; optional fields default to None.
func schemaFieldDecl(field ir.BodyFieldMeta) string {
	pyType := schemaFormatToPython(field.Format)
	if field.Required {
		return fmt.Sprintf("    %s: %s\n", field.Name, pyType)
	}
	return fmt.Sprintf("    %s: Optional[%s] = None\n", field.Name, pyType)
}
