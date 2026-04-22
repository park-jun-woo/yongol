//ff:type feature=contract type=model
//ff:what FuncSignature — Go 함수 시그니처(이름·파라미터·반환·error 여부) 구조체

package contract

// FuncSignature describes the visible contract of the first non-init
// function declaration in a source file. Body-internal details (local
// variables, statements) are deliberately excluded — preserve drift
// must tolerate user rewrites of the body.
type FuncSignature struct {
	Name    string
	Params  []FuncParam
	Returns []string
	HasErr  bool
}
