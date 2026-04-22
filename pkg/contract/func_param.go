//ff:type feature=contract type=model
//ff:what FuncParam — Go 함수 파라미터 이름과 타입 문자열

package contract

// FuncParam represents a single parameter in a Go function signature.
// A parameter declared with multiple names (e.g. `a, b int`) expands
// to one FuncParam per name; an anonymous parameter (common in
// interface-satisfying stubs) yields Name == "".
type FuncParam struct {
	Name string
	Type string
}
