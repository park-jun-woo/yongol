//ff:func feature=gen-gogin type=util control=selection
//ff:what nativeBoolean — BOOLEAN/BOOL 매핑 (NOT NULL → bool, NULL → *bool)

package types

// nativeBoolean returns the binding for a boolean column.
func nativeBoolean(notNull bool, defaultLiteral string) GoTypeBinding {
	switch notNull {
	case false:
		return pointerBoolean(defaultLiteral)
	default:
		return GoTypeBinding{
			SqlcGoType:     "bool",
			ApiField:       "bool",
			ConvertExpr:    "{row}.{field}",
			InsertExpr:     "{var}",
			ResponseExpr:   "{var}.{field}",
			DefaultLiteral: defaultLiteral,
			Kind:           KindNative,
			Supported:      true,
		}
	}
}
