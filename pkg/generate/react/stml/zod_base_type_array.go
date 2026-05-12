//ff:func feature=stml-gen type=util control=sequence
//ff:what array 타입의 item type에서 zod 배열 표현을 반환한다
package stml

import "fmt"

// zodBaseTypeArray returns the zod expression for an array type with
// the given item type. E.g., itemType "string" → "z.array(z.string())".
func zodBaseTypeArray(itemType string) string {
	return fmt.Sprintf("z.array(%s)", zodBaseType(itemType))
}
