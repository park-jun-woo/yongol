//ff:func feature=gen-gogin type=util control=sequence
//ff:what pointerString — NULLABLE string 컬럼의 *string 매핑

package types

// pointerString produces the GoTypeBinding for a nullable string-family
// column.
func pointerString(defaultLiteral string) GoTypeBinding {
	return GoTypeBinding{
		SqlcGoType:     "*string",
		ApiField:       "*string",
		ConvertExpr:    "{row}.{field}",
		InsertExpr:     "{var}",
		ResponseExpr:   "{var}.{field}",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPointer,
		Supported:      true,
	}
}
