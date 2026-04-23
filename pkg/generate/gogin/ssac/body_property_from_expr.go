//ff:func feature=gen-gogin type=util control=sequence
//ff:what bodyPropertyFromExpr — request.<prop> / request.body.<prop> 식에서 최상위 body 속성명 추출
package ssac

import "strings"

// bodyPropertyFromExpr extracts the OpenAPI body property name from an
// SSaC source expression of the form `request.<prop>` or
// `request.body.<prop>`. Returns "" for other shapes so
// maybeMarshalJSONB only triggers on known body-sourced values.
func bodyPropertyFromExpr(expr string) string {
	expr = strings.TrimSpace(expr)
	const reqPrefix = "request."
	if !strings.HasPrefix(expr, reqPrefix) {
		return ""
	}
	rest := strings.TrimPrefix(expr, reqPrefix)
	rest = strings.TrimPrefix(rest, "body.")
	if i := strings.IndexAny(rest, ". "); i >= 0 {
		rest = rest[:i]
	}
	return rest
}
