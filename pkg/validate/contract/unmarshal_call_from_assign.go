//ff:func feature=validate-contract type=util control=sequence topic=preserve-safety
//ff:what unmarshalCallFromAssign — AssignStmt 우변이 json/yaml Unmarshal 호출이면 반환

package contract

import "go/ast"

// unmarshalCallFromAssign returns the underlying CallExpr when the
// assignment's RHS is a single json.Unmarshal / yaml.Unmarshal call
// (with optional v3 package selector tolerated by isUnmarshalCall).
// Otherwise returns nil, signalling the caller to skip.
func unmarshalCallFromAssign(as *ast.AssignStmt) *ast.CallExpr {
	if as == nil || len(as.Rhs) != 1 {
		return nil
	}
	call, ok := as.Rhs[0].(*ast.CallExpr)
	if !ok {
		return nil
	}
	if !isUnmarshalCall(call) {
		return nil
	}
	return call
}
