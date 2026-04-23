//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what scanFileForCurrentUserAssertion — preserved 파일에서 PRV-11 위반 AssignStmt 수집

package contract

import (
	"fmt"
	"go/ast"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// scanFileForCurrentUserAssertion parses path and emits a PRV-11
// Diagnostic for every `ctx.Value("currentUser").(T)` type assertion
// that is NOT written in the comma-ok 2-tuple form. Assignments of
// the shape `cu, ok := ctx.Value(...).(*T)` are silently accepted —
// that is the generated pattern we want to preserve.
func scanFileForCurrentUserAssertion(path string) []diagnostic.Diagnostic {
	fset, file, err := parseGoFile(path)
	if err != nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	ast.Inspect(file, func(n ast.Node) bool {
		assign, ok := n.(*ast.AssignStmt)
		if !ok || len(assign.Rhs) != 1 {
			return true
		}
		if !isCurrentUserAssertion(assign.Rhs[0]) {
			return true
		}
		if len(assign.Lhs) == 2 {
			return true
		}
		line := fset.Position(assign.Pos()).Line
		if hasNolint(fset, file, line, "prv-11") {
			return true
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    path,
			Line:    line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[PRV-11] preserved file uses unsafe single-value currentUser type assertion (line %d)", line),
			Advice: "Use the comma-ok form to avoid nil deref panic:\n" +
				"  cu, ok := ctx.Value(\"currentUser\").(*model.UserClaim)\n" +
				"  if !ok || cu == nil { return api.Error401, nil }",
		})
		return true
	})
	return diags
}
