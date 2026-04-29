//ff:func feature=gen-gogin type=util control=selection
//ff:what nativeString — VARCHAR/TEXT/CHAR/BPCHAR 매핑 (NOT NULL → string, NULL → *string)

package types

// nativeString returns the binding for a string-family column. The
// parameter list (e.g. "255" for VARCHAR(255)) is not encoded into the
// binding itself; length validation lives in the OpenAPI / DDL layer.
func nativeString(notNull bool, defaultLiteral string) GoTypeBinding {
	switch notNull {
	case false:
		return pointerString(defaultLiteral)
	default:
		return GoTypeBinding{
			SqlcGoType:     "string",
			ApiField:       "string",
			ConvertExpr:    "{row}.{field}",
			InsertExpr:     "{var}",
			ResponseExpr:   "{var}.{field}",
			DefaultLiteral: defaultLiteral,
			Kind:           KindNative,
			Supported:      true,
		}
	}
}
