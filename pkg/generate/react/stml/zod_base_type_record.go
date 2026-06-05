//ff:func feature=stml-gen type=util control=sequence
//ff:what object(맵) 타입의 값 타입에서 zod record 표현을 반환한다
package stml

import "fmt"

// zodBaseTypeRecord returns the zod expression for an object(map) type with the
// given additionalProperties value type. E.g., valueType "string" →
// "z.record(z.string())". The free-form marker "*" (or empty) →
// "z.record(z.unknown())".
func zodBaseTypeRecord(valueType string) string {
	if valueType == "" || valueType == "*" {
		return "z.record(z.unknown())"
	}
	return fmt.Sprintf("z.record(%s)", zodBaseType(valueType))
}
