//ff:type feature=ssacmeta type=model
//ff:what Field — params / returns 의 단일 필드

package ssacmeta

// Field is a single named/typed field in params or returns.
type Field struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}
