//ff:type feature=gen-gogin type=model
//ff:what DepthReport — 한 함수의 최대 nesting depth 측정 결과 (filefunc Q1 근거)

package qcheck

// DepthReport carries the maximum control-structure nesting depth observed in
// one function. Depth counts *Stmt nodes that introduce a new control block:
// IfStmt, ForStmt, RangeStmt, SwitchStmt, TypeSwitchStmt, SelectStmt. It
// mirrors filefunc Q1's semantics so generator templates can self-check
// before writing files.
type DepthReport struct {
	Func     string
	MaxDepth int
}
