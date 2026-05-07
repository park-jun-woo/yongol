//ff:func feature=gen-gogin type=util control=selection
//ff:what nativeFloatWithHead — nullable float head 토큰 기반 pgtype Float4/Float8 분기

package types

// nativeFloatWithHead dispatches nullable float to the correct pgtype
// wrapper based on head token (FLOAT4/REAL → Float4, others → Float8).
func nativeFloatWithHead(head string, notNull bool, defaultLiteral string) GoTypeBinding {
	if notNull {
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
	switch head {
	case "REAL", "FLOAT4":
		return pgtypeFloat4(defaultLiteral)
	default:
		return pgtypeFloat8(defaultLiteral)
	}
}
