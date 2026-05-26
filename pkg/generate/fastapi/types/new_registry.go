//ff:func feature=gen-fastapi type=util control=sequence
//ff:what NewRegistry — FastAPITypeRegistry 팩토리 생성자

package types

// NewRegistry returns a new FastAPITypeRegistry.
func NewRegistry() *FastAPITypeRegistry {
	return &FastAPITypeRegistry{}
}
