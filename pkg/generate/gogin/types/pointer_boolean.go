//ff:func feature=gen-gogin type=util control=sequence
//ff:what pointerBoolean — NULLABLE boolean 컬럼의 *bool 매핑

package types

// pointerBoolean produces the GoTypeBinding for a nullable boolean
// column.
func pointerBoolean(defaultLiteral string) GoTypeBinding {
	return GoTypeBinding{
		SqlcGoType:     "*bool",
		ApiField:       "*bool",
		ConvertExpr:    "{row}.{field}",
		InsertExpr:     "{var}",
		ResponseExpr:   "{var}.{field}",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPointer,
		Supported:      true,
	}
}
