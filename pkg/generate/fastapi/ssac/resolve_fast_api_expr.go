//ff:func feature=gen-fastapi type=util control=sequence
//ff:what resolveFastAPIExpr — raw SSaC 표현식을 FastAPI 용으로 변환 (request.X → body.x)

package ssac

import "strings"

// resolveFastAPIExpr rewrites raw SSaC expressions for FastAPI. In particular,
// "request.Xxx" is mapped to "body.xxx" (snake_case) because FastAPI service
// methods receive the request body as a separate "body" parameter.
func resolveFastAPIExpr(expr string) string {
	if strings.HasPrefix(expr, "request.") {
		field := expr[len("request."):]
		return "body." + snakeCase(field)
	}
	return expr
}
