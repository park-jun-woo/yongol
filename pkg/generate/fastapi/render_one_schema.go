//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what renderOneSchema — 단일 plan 의 요청 본문에 대한 Pydantic BaseModel 클래스 출력

package fastapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// renderOneSchema writes a single Pydantic BaseModel class for a plan's
// request body.
func renderOneSchema(b *strings.Builder, plan *ir.ServicePlan) {
	className := schemaPascalCase(plan.OperationID) + "Request"
	b.WriteString(fmt.Sprintf("class %s(BaseModel):\n", className))
	for _, field := range plan.BodyFields {
		b.WriteString(schemaFieldDecl(field))
	}
	b.WriteString("\n\n")
}
