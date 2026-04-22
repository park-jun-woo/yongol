//ff:func feature=gen-splitter type=util control=sequence
//ff:what methodReceiver — FuncDecl 의 receiver 타입 이름 추출 (method 아니면 "")
package splitter

import "go/ast"

// methodReceiver returns the receiver type name for a FuncDecl, or ""
// when the decl is a plain (non-method) function. It delegates to
// receiverName for the T / *T pattern match.
func methodReceiver(d *ast.FuncDecl) string {
	if d.Recv == nil || len(d.Recv.List) == 0 {
		return ""
	}
	return receiverName(d.Recv.List[0].Type)
}
