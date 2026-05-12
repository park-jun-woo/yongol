//ff:type feature=design-parse type=model
//ff:what TypographyToken — 단일 typography 디자인 토큰 구조체

package design

// TypographyToken represents a single typography design token.
type TypographyToken struct {
	FontFamily    string `yaml:"fontFamily"`
	FontSize      string `yaml:"fontSize"`
	FontWeight    string `yaml:"fontWeight"`
	LineHeight    string `yaml:"lineHeight"`
	LetterSpacing string `yaml:"letterSpacing"`
}
