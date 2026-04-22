//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what scanFileForNilDeref — preserved 파일에서 Get/Find 반환값 즉시 selector 접근 진단

package contract

import (
	"fmt"
	"go/ast"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// scanFileForNilDeref looks for the chained shape
// `<recv>.GetSomething().Field` or `<recv>.FindSomething().Field`.
// Such lookups commonly return `(*T, nil)` → `(nil, err)` and
// dereferencing without a guard is the exact bug PRV-16 targets.
// Free-function calls (no receiver) are skipped — they are rarer and
// add too many false positives on pure utilities like `time.Now()`.
func scanFileForNilDeref(path string) []diagnostic.Diagnostic {
	fset, file, err := parseGoFile(path)
	if err != nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	ast.Inspect(file, func(n ast.Node) bool {
		sel, ok := n.(*ast.SelectorExpr)
		if !ok {
			return true
		}
		call, ok := sel.X.(*ast.CallExpr)
		if !ok {
			return true
		}
		name := callMethodName(call)
		if !strings.HasPrefix(name, "Get") && !strings.HasPrefix(name, "Find") {
			return true
		}
		line := fset.Position(sel.Pos()).Line
		if hasNolint(fset, file, line, "prv-16") {
			return true
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    path,
			Line:    line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[PRV-16] preserved file dereferences %s() return value without nil guard (line %d)", name, line),
			Advice: "Bind the result and guard before access:\n" +
				"  v := recv." + name + "()\n" +
				"  if v == nil { return api.Error404, nil }\n" +
				"  _ = v.Field",
		})
		return true
	})
	return diags
}
