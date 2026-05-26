//ff:func feature=gen-nestjs type=util control=sequence
//ff:what NewRegistry — NestJSTypeRegistry 팩토리 생성자

package types

// NewRegistry returns a new NestJSTypeRegistry.
func NewRegistry() *NestJSTypeRegistry {
	return &NestJSTypeRegistry{}
}
