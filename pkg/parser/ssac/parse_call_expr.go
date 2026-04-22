//ff:func feature=ssac-parse type=parser control=sequence
//ff:what "Model.Method(args)" 호출 표현식 파싱
package ssac

import "strings"

// parseCallExpr는 "Model.Method(args)" 또는 "pkg.Func(args)"를 파싱한다.
func parseCallExpr(expr string) (string, []Arg) {
	expr = strings.TrimSpace(expr)
	parenIdx := strings.Index(expr, "(")
	if parenIdx < 0 {
		return expr, nil
	}
	model := expr[:parenIdx]
	argsStr := expr[parenIdx+1:]
	argsStr = strings.TrimSuffix(strings.TrimSpace(argsStr), ")")
	argsStr = strings.TrimSpace(argsStr)
	if argsStr == "" {
		return model, nil
	}
	return model, parseArgs(argsStr)
}
