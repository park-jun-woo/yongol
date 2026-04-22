//ff:type feature=funcspec type=model
//ff:what 구조체 필드 정보 타입
package funcspec

// Field represents a struct field.
type Field struct {
	Name     string
	Type     string
	JSONName string // json tag name (empty = use Name)
}
