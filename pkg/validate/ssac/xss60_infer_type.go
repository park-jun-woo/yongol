//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what xss60InferType — publish payload 소스 표현식의 Go 타입 추론 (literal/dot expression)

package ssac

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// xss60InferType infers the Go type of a publish payload source expression.
func xss60InferType(expr string, fn parsessac.ServiceFunc, tableMap map[string]*ddl.Table) string {
	expr = strings.TrimSpace(expr)
	if expr == "" {
		return ""
	}

	// Case 1: string literal ("completed", "hello")
	if len(expr) >= 2 && expr[0] == '"' && expr[len(expr)-1] == '"' {
		return "string"
	}

	// Case 2: integer literal (0, 42, 86400)
	if xss60IsIntLiteral(expr) {
		return "int64"
	}

	// Case 3: dot expression (wf.ID, user.Email)
	if idx := strings.IndexByte(expr, '.'); idx > 0 && idx < len(expr)-1 {
		return xss60InferDotExpr(expr, idx, fn, tableMap)
	}

	// Case 4: cannot infer
	return ""
}
