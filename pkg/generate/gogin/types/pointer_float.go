//ff:func feature=gen-gogin type=util control=sequence
//ff:what pointerFloat — NULLABLE float 컬럼의 *float64 매핑

package types

// pointerFloat produces the GoTypeBinding for a nullable float-family
// column.
func pointerFloat(defaultLiteral string) GoTypeBinding {
	return GoTypeBinding{
		SqlcGoType:     "*float64",
		ApiField:       "*float64",
		ConvertExpr:    "{row}.{field}",
		InsertExpr:     "{var}",
		ResponseExpr:   "{var}.{field}",
		DefaultLiteral: defaultLiteral,
		Kind:           KindPointer,
		Supported:      true,
	}
}
