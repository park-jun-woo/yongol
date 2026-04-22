//ff:func feature=funcspec type=util control=sequence
//ff:what 함수 본문이 스텁(빈 본문 / panic / zero-only return)인지 확인한다
package funcspec

import (
	"go/ast"
	"go/token"
)

// isStubBody returns true only for genuinely empty implementations:
//   - empty body
//   - single panic(...) call (typically panic("TODO"))
//   - single return whose every result is a zero literal / nil / empty
//     composite literal (e.g. `return Resp{}, nil`)
//
// A single return that carries any meaningful value
// (`return Resp{Status: "ok"}, nil`) is treated as a real implementation.
func isStubBody(fset *token.FileSet, body *ast.BlockStmt) bool {
	if len(body.List) == 0 {
		return true
	}
	if len(body.List) > 1 {
		return false
	}
	if isPanicCall(body.List[0]) {
		return true
	}
	if ret, ok := body.List[0].(*ast.ReturnStmt); ok {
		return allZeroResults(ret.Results)
	}
	return false
}
