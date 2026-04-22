//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what scanFileForMapAccess — preserved 파일에서 가드 없는 map[key].Sel 패턴 진단 수집

package contract

import (
	"fmt"
	"go/ast"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// scanFileForMapAccess walks the file AST looking for SelectorExpr
// whose base is an IndexExpr. An index-then-select chain is the shape
// most likely to deref a nil pointer value stored in a map — without
// a comma-ok guard ahead of it, the author has no way to know the key
// is present. We only flag the INLINE form because separate-variable
// assignments (`u := m[k]; u.X`) require type reasoning that our
// parser-only pass cannot provide.
func scanFileForMapAccess(path string) []diagnostic.Diagnostic {
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
		idx, ok := sel.X.(*ast.IndexExpr)
		if !ok {
			return true
		}
		if leftmostIdentName(idx.X) == "" {
			return true
		}
		line := fset.Position(sel.Pos()).Line
		if hasNolint(fset, file, line, "prv-15") {
			return true
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    path,
			Line:    line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[PRV-15] preserved file dereferences map/slice indexed value without comma-ok guard (line %d)", line),
			Advice: "Split the access and guard:\n" +
				"  v, ok := m[k]\n" +
				"  if !ok || v == nil { return api.Error404, nil }\n" +
				"  _ = v.Field",
		})
		return true
	})
	return diags
}
