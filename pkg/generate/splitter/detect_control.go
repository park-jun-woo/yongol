//ff:func feature=gen-splitter type=util control=iteration dimension=1
//ff:what detectControl — 함수 body AST 의 depth-1 제어 구조를 감지해 control= 값 결정
package splitter

import "go/ast"

// detectControl inspects the direct children of body and classifies the
// function's control shape per Böhm–Jacopini:
//
//   - iteration: depth-1 for/range exists
//   - selection: depth-1 switch/type-switch exists (and no loop)
//   - sequence : otherwise
//
// When iteration is returned, dimension is always "1" for splitter output —
// external tools do not emit nested iteration on splitter-managed funcs.
// If both loop and switch appear at depth 1 the body is too complex to
// re-annotate safely, so detectControl falls back to sequence and the
// caller should treat that as a signal to not split further.
func detectControl(body *ast.BlockStmt) (control, dimension string) {
	if body == nil {
		return "sequence", ""
	}
	hasLoop := false
	hasSwitch := false
	for _, stmt := range body.List {
		switch stmt.(type) {
		case *ast.ForStmt, *ast.RangeStmt:
			hasLoop = true
		case *ast.SwitchStmt, *ast.TypeSwitchStmt:
			hasSwitch = true
		}
	}
	if hasLoop && !hasSwitch {
		return "iteration", "1"
	}
	if hasSwitch && !hasLoop {
		return "selection", ""
	}
	return "sequence", ""
}
