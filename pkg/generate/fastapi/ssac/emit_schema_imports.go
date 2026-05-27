//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what emitSchemaImports — Pydantic 스키마 클래스 import 출력

package ssac

import (
	"fmt"
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// emitSchemaImports writes Pydantic schema imports for service function
// parameters that reference body models. For example, a login service
// with LoginRequest body parameter needs:
//
//	from app.schemas.auth import LoginRequest
func emitSchemaImports(b *strings.Builder, plans []*ir.ServicePlan, feature string) {
	var models []string
	seen := make(map[string]bool)
	for _, plan := range plans {
		method := strings.ToUpper(plan.HTTPMethod)
		hasBody := (method == "POST" || method == "PUT" || method == "PATCH") && len(plan.BodyFields) > 0
		if !hasBody {
			continue
		}
		name := pascalCase(plan.OperationID) + "Request"
		if !seen[name] {
			seen[name] = true
			models = append(models, name)
		}
	}
	if len(models) == 0 {
		return
	}
	sort.Strings(models)
	b.WriteString(fmt.Sprintf("from app.schemas.%s import %s\n", feature, strings.Join(models, ", ")))
}
