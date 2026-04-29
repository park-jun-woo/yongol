//ff:func feature=gen-gogin type=util control=selection
//ff:what nativeInteger — BIGINT/INT/SMALLINT 계열 매핑 (NOT NULL → int64, NULL → *int64)

package types

// nativeInteger returns the binding for an integer-family column.
// notNull selects the NOT NULL (Native) vs NULLABLE (Pointer) form.
// defaultLiteral is forwarded into GoTypeBinding.DefaultLiteral verbatim.
func nativeInteger(notNull bool, defaultLiteral string) GoTypeBinding {
	switch notNull {
	case true:
		return GoTypeBinding{
			SqlcGoType:     "int64",
			ApiField:       "int64",
			ConvertExpr:    "{row}.{field}",
			InsertExpr:     "{var}",
			ResponseExpr:   "{var}.{field}",
			DefaultLiteral: defaultLiteral,
			Kind:           KindNative,
			Supported:      true,
		}
	default:
		return pointerInteger(defaultLiteral)
	}
}
