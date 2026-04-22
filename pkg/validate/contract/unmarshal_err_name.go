//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what unmarshalErrName — AssignStmt LHS 에서 error 를 받는 식별자 이름 추출

package contract

import (
	"go/ast"
	"strings"
)

// unmarshalErrName inspects as.Lhs for the identifier that receives the
// Unmarshal call's error return. Since Unmarshal has a single return
// value, the LHS must be length 1; we reject blank ident ("_") because
// that is an explicit discard (handled elsewhere as "discarded").
//
// The name must be "err" or end with "err"/"Err" — heuristic alignment
// with the rest of the codebase where variant names like `uErr` are
// still meaningful error receivers.
func unmarshalErrName(as *ast.AssignStmt) string {
	if as == nil || len(as.Lhs) != 1 {
		return ""
	}
	ident, ok := as.Lhs[0].(*ast.Ident)
	if !ok {
		return ""
	}
	if ident.Name == "_" {
		return ""
	}
	lower := strings.ToLower(ident.Name)
	if lower == "err" || strings.HasSuffix(lower, "err") {
		return ident.Name
	}
	return ""
}
