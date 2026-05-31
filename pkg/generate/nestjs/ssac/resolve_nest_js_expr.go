//ff:func feature=gen-nestjs type=util control=sequence
//ff:what resolveNestJSExpr — raw SSaC 표현식을 NestJS 용으로 변환 (request.x → body.x)

package ssac

import "strings"

// resolveNestJSExpr rewrites raw SSaC expressions for NestJS. In particular,
// "request.xxx" is mapped to "body.xxx" because NestJS service methods receive
// the request body as a separate "body" parameter.
func resolveNestJSExpr(expr string) string {
	if strings.HasPrefix(expr, "request.") {
		return "body." + expr[len("request."):]
	}
	return expr
}
