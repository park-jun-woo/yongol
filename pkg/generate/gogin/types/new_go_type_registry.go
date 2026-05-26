//ff:func feature=gen-gogin type=util control=sequence
//ff:what NewGoTypeRegistry — GoTypeRegistry 팩토리 생성자

package types

// NewGoTypeRegistry returns a new GoTypeRegistry.
func NewGoTypeRegistry() *GoTypeRegistry {
	return &GoTypeRegistry{}
}
