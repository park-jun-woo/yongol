//ff:func feature=stml-gen type=util control=selection
//ff:what 타입 문자열에서 zod 기본 타입 표현을 반환한다
package stml

// zodBaseType returns the base zod type expression for a given scalar type
// string. Allowed: "string"/""/"integer"/"number"/"boolean". array/object/enum
// are handled by zodChain before reaching here. An unknown type panics with a
// *zodGenError (recovered at the GenerateWith boundary) instead of silently
// downgrading to z.string().
func zodBaseType(typ string) string {
	switch typ {
	case "string", "":
		return "z.string()"
	case "integer":
		return "z.number().int()"
	case "number":
		return "z.number()"
	case "boolean":
		return "z.boolean()"
	default:
		panic(&zodGenError{Type: typ})
	}
}
