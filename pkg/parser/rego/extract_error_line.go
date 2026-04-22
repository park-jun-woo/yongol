//ff:func feature=orchestrator type=util control=sequence
//ff:what opa/ast 에러에서 라인 번호를 추출한다
package rego

import "github.com/open-policy-agent/opa/ast"

// extractErrorLine extracts the line number from an OPA parse error.
// Returns 0 if the error does not contain location info.
func extractErrorLine(err error) int {
	errs, ok := err.(ast.Errors)
	if !ok || len(errs) == 0 || errs[0].Location == nil {
		return 0
	}
	return errs[0].Location.Row
}
