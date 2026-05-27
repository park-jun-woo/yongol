//ff:func feature=gen-fastapi type=util control=sequence
//ff:what writeHTTPHandler — ServicePlan 메타데이터 기반 FastAPI HTTP route handler 함수 작성

package ssac

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

// writeHTTPHandler writes a decorated FastAPI route handler function.
// Parameter declarations are driven by ServicePlan metadata:
//   - PathParams → typed path parameters as function args
//   - BodyFields → Pydantic model parameter
//   - QueryParams → typed query parameters as function args
func writeHTTPHandler(b *strings.Builder, plan *ir.ServicePlan) {
	decorator := pyHTTPDecorator(plan.HTTPMethod)
	routePath := routeSuffix(plan)
	funcName := snakeCase(plan.OperationID)

	method := strings.ToUpper(plan.HTTPMethod)
	hasBody := (method == "POST" || method == "PUT" || method == "PATCH") && len(plan.BodyFields) > 0
	hasQuery := len(plan.QueryParams) > 0
	hasPath := len(plan.PathParams) > 0

	b.WriteString(fmt.Sprintf("@router.%s(\"%s\")\n", decorator, routePath))
	b.WriteString(fmt.Sprintf("async def %s(\n", funcName))

	// Path parameters as typed function arguments.
	for _, pp := range plan.PathParams {
		b.WriteString(fmt.Sprintf("    %s: int,\n", pp))
	}

	// Body parameter (Pydantic model).
	if hasBody {
		reqModel := pascalCase(plan.OperationID) + "Request"
		b.WriteString(fmt.Sprintf("    body: %s,\n", reqModel))
	}

	// Query parameters as typed function arguments.
	for _, qp := range plan.QueryParams {
		pyType := openAPITypeToPython(qp.Type)
		if qp.Required {
			b.WriteString(fmt.Sprintf("    %s: %s,\n", qp.Name, pyType))
		} else {
			b.WriteString(fmt.Sprintf("    %s: %s | None = None,\n", qp.Name, pyType))
		}
	}

	// Dependency injections.
	b.WriteString("    current_user: dict = Depends(get_current_user),\n")
	b.WriteString("    session: AsyncSession = Depends(get_session),\n")
	b.WriteString("):\n")

	// Build service call arguments.
	var callArgs []string
	callArgs = append(callArgs, "session")
	if hasPath {
		for _, pp := range plan.PathParams {
			callArgs = append(callArgs, pp)
		}
	}
	if hasBody {
		callArgs = append(callArgs, "body")
	}
	if hasQuery {
		for _, qp := range plan.QueryParams {
			callArgs = append(callArgs, qp.Name)
		}
	}
	callArgs = append(callArgs, "current_user")

	b.WriteString(fmt.Sprintf("    return await svc.%s(%s)\n",
		funcName, strings.Join(callArgs, ", ")))
}

// openAPITypeToPython maps an OpenAPI type string to a Python type annotation.
func openAPITypeToPython(t string) string {
	switch t {
	case "integer":
		return "int"
	case "number":
		return "float"
	case "boolean":
		return "bool"
	default:
		return "str"
	}
}
