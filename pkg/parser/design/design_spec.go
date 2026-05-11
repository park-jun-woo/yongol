//ff:type feature=frontend type=model
//ff:what DESIGN.md 파싱 결과를 담는 디자인 토큰 구조체
package design

// DesignSpec holds parsed design tokens from a DESIGN.md file.
// YAML front matter provides token maps; Markdown body yields Headings.
type DesignSpec struct {
	File       string
	Version    string                   `yaml:"version"`
	Name       string                   `yaml:"name"`
	Colors     map[string]string        `yaml:"colors"`
	Typography map[string]TypographyToken `yaml:"typography"`
	Rounded    map[string]string        `yaml:"rounded"`
	Spacing    map[string]string        `yaml:"spacing"`
	Components map[string]ComponentToken `yaml:"components"`
	Headings   []string                 // ## headings extracted from body
}

// TypographyToken represents a single typography design token.
type TypographyToken struct {
	FontFamily    string `yaml:"fontFamily"`
	FontSize      string `yaml:"fontSize"`
	FontWeight    string `yaml:"fontWeight"`
	LineHeight    string `yaml:"lineHeight"`
	LetterSpacing string `yaml:"letterSpacing"`
}

// ComponentToken represents a component-level design token.
type ComponentToken struct {
	Props map[string]string `yaml:"props"`
}
