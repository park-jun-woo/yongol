//ff:func feature=validate-contract type=util control=selection topic=preserve-safety
//ff:what isUnmarshalCall — call 이 json/yaml(.v2/.v3) Unmarshal 호출인지 판정

package contract

import "go/ast"

// isUnmarshalCall recognises the `<pkg>.Unmarshal(...)` shape where
// pkg is one of the common serialization packages. We look at the
// leaf selector name only — aliasing (e.g. `jsoniter` imported as
// `json`) is still caught because the source text uses the bound
// name, which is what the AST carries.
func isUnmarshalCall(call *ast.CallExpr) bool {
	if call == nil {
		return false
	}
	sel, ok := call.Fun.(*ast.SelectorExpr)
	if !ok || sel.Sel == nil || sel.Sel.Name != "Unmarshal" {
		return false
	}
	pkg, ok := sel.X.(*ast.Ident)
	if !ok {
		return false
	}
	switch pkg.Name {
	case "json", "yaml", "toml", "xml":
		return true
	}
	return false
}
