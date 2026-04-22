//ff:func feature=ssac-parse type=parser control=sequence
//ff:what parseCallExprInputs — parses the "pkg.Func({Key: val, ...}) remainder" form
package ssac

import "strings"

// parseCallExprInputs parses "pkg.Func({Key: val, ...}) remainder".
// Also returns the remainder after the closing parenthesis.
func parseCallExprInputs(expr string) (string, map[string]string, string, error) {
	expr = strings.TrimSpace(expr)
	parenIdx := strings.Index(expr, "(")
	if parenIdx < 0 {
		return expr, nil, "", nil
	}
	model := expr[:parenIdx]
	afterParen := expr[parenIdx+1:]
	// find the last )
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
