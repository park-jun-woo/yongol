//ff:func feature=gen-gogin type=util control=selection
//ff:what composeArrayBinding — array element 분기 (지원 → []T 바인딩, 미지원 → unsupportedBinding)

package types

// composeArrayBinding builds the GoTypeBinding for an array column once
// the element type has been resolved. Split out from arrayBinding so the
// dispatch (selection) and the literal construction (sequence) live in
// separate funcs per filefunc F1 / A10.
func composeArrayBinding(goElem string, supported bool, elementHead, defaultLiteral string) GoTypeBinding {
	switch supported {
	case false:
		return unsupportedBinding("array element type " + elementHead + " is not supported")
	}
	sliceType := "[]" + goElem
	return GoTypeBinding{
		SqlcGoType:     sliceType,
		ApiField:       sliceType,
		ConvertExpr:    "{row}.{field}",
		InsertExpr:     "{var}",
		ResponseExpr:   "{var}.{field}",
		DefaultLiteral: defaultLiteral,
		Kind:           KindArray,
		Supported:      true,
	}
}
