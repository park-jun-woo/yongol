//ff:type feature=projectconfig type=model
//ff:what 프론트엔드 설정 구조체 (Theme 은 tailwind.config 방출용 semantic 토큰)
package manifest

// Frontend holds the React/TSX scaffold configuration. Theme maps directly
// to the tailwind.config.js `theme.extend.colors` block that yongol emits;
// absent fields fall back to shadcn defaults.
type Frontend struct {
	Lang      string         `yaml:"lang"`
	Framework string         `yaml:"framework"`
	Bundler   string         `yaml:"bundler"`
	Name      string         `yaml:"name"`
	Theme     *FrontendTheme `yaml:"theme,omitempty"`
}

// FrontendTheme mirrors shadcn/ui's canonical semantic tokens. All fields
// are optional; unset tokens inherit shadcn defaults in the generated
// tailwind.config.js. Colors accept any CSS color string (hex, rgb, hsl).
type FrontendTheme struct {
	Primary     string `yaml:"primary,omitempty"`
	Secondary   string `yaml:"secondary,omitempty"`
	Accent      string `yaml:"accent,omitempty"`
	Destructive string `yaml:"destructive,omitempty"`
	Muted       string `yaml:"muted,omitempty"`
	Background  string `yaml:"background,omitempty"`
	Foreground  string `yaml:"foreground,omitempty"`
	Border      string `yaml:"border,omitempty"`
	Radius      string `yaml:"radius,omitempty"` // e.g. "0.5rem"
}
