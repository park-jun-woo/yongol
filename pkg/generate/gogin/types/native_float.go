//ff:func feature=gen-gogin type=util control=selection
//ff:what nativeFloat — REAL/FLOAT/FLOAT4/FLOAT8 매핑 (NOT NULL → float64, NULL → *float64)

package types

// nativeFloat returns the binding for a float-family column.
func nativeFloat(notNull bool, defaultLiteral string) GoTypeBinding {
	switch notNull {
	case false:
		return pointerFloat(defaultLiteral)
	default:
		return GoTypeBinding{
			SqlcGoType:     "float64",
			ApiField:       "float64",
			ConvertExpr:    "{row}.{field}",
			InsertExpr:     "{var}",
			ResponseExpr:   "{var}.{field}",
			DefaultLiteral: defaultLiteral,
			Kind:           KindNative,
			Supported:      true,
		}
	}
}
