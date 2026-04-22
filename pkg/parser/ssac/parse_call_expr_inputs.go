//ff:func feature=ssac-parse type=parser control=sequence
//ff:what "pkg.Func({Key: val, ...}) remainder" 형식 파싱
package ssac

import "strings"

// parseCallExprInputs는 "pkg.Func({Key: val, ...}) remainder"를 파싱한다.
// 닫는 괄호 뒤의 remainder도 반환한다.
func parseCallExprInputs(expr string) (string, map[string]string, string, error) {
	expr = strings.TrimSpace(expr)
	parenIdx := strings.Index(expr, "(")
	if parenIdx < 0 {
		return expr, nil, "", nil
	}
	model := expr[:parenIdx]
	afterParen := expr[parenIdx+1:]
	// 마지막 ) 찾기
	closeIdx := strings.LastIndex(afterParen, ")")
	if closeIdx < 0 {
		return model, nil, "", nil
	}
	inner := strings.TrimSpace(afterParen[:closeIdx])
	remainder := strings.TrimSpace(afterParen[closeIdx+1:])
	if inner == "" {
		return model, nil, remainder, nil
	}
	inputs, err := parseInputs(inner)
	return model, inputs, remainder, err
}
