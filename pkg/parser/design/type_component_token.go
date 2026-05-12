//ff:type feature=design-parse type=model
//ff:what ComponentToken — 컴포넌트 레벨 디자인 토큰 구조체

package design

// ComponentToken represents a component-level design token.
type ComponentToken struct {
	Base           string            `yaml:"base"`
	Variants       map[string]string `yaml:"variants"`
	Sizes          map[string]string `yaml:"sizes"`
	DefaultVariant string            `yaml:"defaultVariant"`
	DefaultSize    string            `yaml:"defaultSize"`
	Props          map[string]string `yaml:"props"`
}
